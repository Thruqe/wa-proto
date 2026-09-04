package extractor

import (
	"testing"
)

func TestExtractor_ProcessSource(t *testing.T) {
	sampleJS := `
	__d("WAAdv", [], (function(a, b, c, d, e, f) {
		const ADVEncryptionType = {
			E2EE: 0,
			HOSTED: 1,
			NON_E2EE: 2
		};
		f.ADVEncryptionType = ADVEncryptionType;

		f.ADVKeyIndexListSpec = {
			rawId: [1, b.TYPES.UINT32],
			timestamp: [2, b.TYPES.UINT64],
			currentIndex: [3, b.TYPES.UINT32],
			validIndexes: [4, b.FLAGS.PACKED | b.TYPES.UINT32 | b.FLAGS.REPEATED],
			accountType: [5, b.TYPES.ENUM, ADVEncryptionType],
			__oneofs__: {
				choice: ["rawId", "timestamp"]
			}
		};
	}));
	`

	ext := NewExtractor("2.3000.1000")
	if err := ext.ProcessSource(sampleJS); err != nil {
		t.Fatalf("ProcessSource failed: %v", err)
	}

	mod := ext.Modules["WAAdv"]
	if mod == nil {
		t.Fatalf("expected WAAdv module, got nil")
	}

	if len(mod.Enums) != 1 {
		t.Errorf("expected 1 enum, got %d", len(mod.Enums))
	}
	if len(mod.Messages) != 1 {
		t.Errorf("expected 1 message, got %d", len(mod.Messages))
	}

	msg := mod.Messages[0]
	if msg.Name != "ADVKeyIndexList" {
		t.Errorf("expected message name ADVKeyIndexList, got %s", msg.Name)
	}
	if len(msg.Oneofs) != 1 {
		t.Errorf("expected 1 oneof, got %d", len(msg.Oneofs))
	}
	if msg.Oneofs[0].Name != "choice" || len(msg.Oneofs[0].Fields) != 2 {
		t.Errorf("expected oneof 'choice' with 2 fields, got %+v", msg.Oneofs[0])
	}
}
