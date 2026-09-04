package corrector

import (
	"strings"
	"testing"

	"github.com/Thruqe/wa-proto/pkg/ast"
)

func TestFixProto3_Message(t *testing.T) {
	schema := &ast.ProtoSchema{
		Syntax: "proto3",
		Messages: []*ast.MessageDef{
			{
				Name: "TestMessage",
				Fields: []*ast.FieldDef{
					{Name: "id", ID: 1, Type: "string", Rule: "required"},
					{Name: "count", ID: 2, Type: "int32", Rule: "optional"},
				},
			},
		},
	}

	FixProto3(schema)

	if schema.Messages[0].Fields[0].Rule != "optional" {
		t.Errorf("expected required to be converted to optional, got %s", schema.Messages[0].Fields[0].Rule)
	}
}

func TestFixProto3_EnumFirstValue(t *testing.T) {
	schema := &ast.ProtoSchema{
		Syntax: "proto3",
		Enums: []*ast.EnumDef{
			{
				Name: "ReminderAction",
				Values: []*ast.EnumValueDef{
					{Name: "NOTIFY", ID: 1},
					{Name: "SNOOZE", ID: 2},
				},
			},
		},
	}

	FixProto3(schema)

	enum := schema.Enums[0]
	if len(enum.Values) != 3 {
		t.Fatalf("expected 3 values, got %d", len(enum.Values))
	}
	if enum.Values[0].ID != 0 {
		t.Errorf("expected first enum value to be 0, got %d", enum.Values[0].ID)
	}
	if enum.Values[0].Name != "REMINDERACTION_UNKNOWN" {
		t.Errorf("expected REMINDERACTION_UNKNOWN, got %s", enum.Values[0].Name)
	}
}

func TestFixProto3Content_String(t *testing.T) {
	raw := `syntax = "proto3";
package waproto;

message Foo {
    required string name = 1;
}

enum Status {
    ACTIVE = 1;
    INACTIVE = 2;
}
`
	fixed := FixProto3Content(raw)
	if strings.Contains(fixed, "required ") {
		t.Errorf("did not remove required")
	}
	if !strings.Contains(fixed, "STATUS_UNKNOWN = 0;") {
		t.Errorf("did not inject zero value enum constant, got:\n%s", fixed)
	}
}
