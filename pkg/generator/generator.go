package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/Thruqe/wa-proto/pkg/ast"
	"github.com/Thruqe/wa-proto/pkg/catalog"
	"github.com/Thruqe/wa-proto/pkg/corrector"
)

// Report contains statistics about the generation process.
type Report struct {
	TotalPackages  int
	TotalMessages  int
	TotalEnums     int
	GeneratedFiles []string
}

// GenerateMonolithic generates a single monolithic WAProto.proto matching wppconnect wa-proto.
func GenerateMonolithic(schema *ast.ProtoSchema, w io.Writer) error {
	// Ensure strict proto3 compliance (no required fields, first enum is 0)
	corrector.FixProto3(schema)

	var b strings.Builder
	b.WriteString("syntax = \"proto3\";\n")
	b.WriteString("package waproto;\n\n")

	if schema.Version != "" {
		b.WriteString(fmt.Sprintf("/// WhatsApp Version: %s\n\n", schema.Version))
	}

	// Sort entities alphabetically for deterministic output
	type entity struct {
		name    string
		content string
	}
	var entities []entity

	for _, e := range schema.Enums {
		entities = append(entities, entity{name: e.Name, content: e.FormatProto2("")})
	}
	for _, m := range schema.Messages {
		entities = append(entities, entity{name: m.Name, content: m.FormatProto2("")})
	}

	sort.Slice(entities, func(i, j int) bool {
		return entities[i].name < entities[j].name
	})

	for _, ent := range entities {
		b.WriteString(ent.content)
		b.WriteString("\n\n")
	}

	_, err := io.WriteString(w, b.String())
	return err
}

// GenerateModular partitions the schema into the modular wa-core/proto packages.
func GenerateModular(schema *ast.ProtoSchema, outDir string) (*Report, error) {
	pkgMessages := make(map[string][]*ast.MessageDef)
	pkgEnums := make(map[string][]*ast.EnumDef)

	// Group enums into packages
	for _, e := range schema.Enums {
		pkg := catalog.LookupPackageForType(e.Name)
		pkgEnums[pkg] = append(pkgEnums[pkg], e)
	}

	// Group messages into packages
	for _, m := range schema.Messages {
		pkg := catalog.LookupPackageForType(m.Name)
		pkgMessages[pkg] = append(pkgMessages[pkg], m)
	}

	// Also ensure all registered catalog packages exist
	allPkgs := make(map[string]bool)
	for p := range catalog.Packages {
		allPkgs[p] = true
	}
	for p := range pkgMessages {
		allPkgs[p] = true
	}
	for p := range pkgEnums {
		allPkgs[p] = true
	}

	var pkgNames []string
	for p := range allPkgs {
		pkgNames = append(pkgNames, p)
	}
	sort.Strings(pkgNames)

	report := &Report{
		TotalPackages: len(pkgNames),
	}

	for _, pkgDir := range pkgNames {
		meta, exists := catalog.Packages[pkgDir]
		if !exists {
			meta = &catalog.PackageMeta{
				Dir:       pkgDir,
				File:      pkgDir + ".proto",
				Package:   pkgDir,
				GoPackage: "go.mau.fi/whatsmeow/proto/" + pkgDir,
			}
		}

		msgs := pkgMessages[pkgDir]
		enums := pkgEnums[pkgDir]

		// Skip empty packages if they have no messages or enums
		if len(msgs) == 0 && len(enums) == 0 {
			continue
		}

		targetDirPath := filepath.Join(outDir, pkgDir)
		if err := os.MkdirAll(targetDirPath, 0755); err != nil {
			return nil, fmt.Errorf("failed creating package dir %s: %w", targetDirPath, err)
		}

		targetFilePath := filepath.Join(targetDirPath, meta.File)

		// Collect referenced types to compute needed imports
		referencedTypes := make(map[string]bool)
		for _, m := range msgs {
			collectMessageTypes(m, referencedTypes)
		}

		// Calculate required imports
		neededImports := make(map[string]bool)
		// Include baseline default imports for this package
		for _, imp := range meta.Imports {
			neededImports[imp] = true
		}

		for t := range referencedTypes {
			refPkg := catalog.LookupPackageForType(t)
			if refPkg != "" && refPkg != pkgDir {
				if refMeta, ok := catalog.Packages[refPkg]; ok {
					impPath := refMeta.Dir + "/" + refMeta.File
					neededImports[impPath] = true
				}
			}
		}

		// Format the file
		var b strings.Builder
		b.WriteString("syntax = \"proto2\";\n")
		b.WriteString(fmt.Sprintf("package %s;\n", meta.Package))
		b.WriteString(fmt.Sprintf("option go_package = %q;\n", meta.GoPackage))

		if len(neededImports) > 0 {
			b.WriteString("\n")
			var imps []string
			for imp := range neededImports {
				imps = append(imps, imp)
			}
			sort.Strings(imps)
			for _, imp := range imps {
				b.WriteString(fmt.Sprintf("import %q;\n", imp))
			}
		}

		b.WriteString("\n")

		// Sort and write Enums
		sort.Slice(enums, func(i, j int) bool {
			return enums[i].Name < enums[j].Name
		})
		for _, e := range enums {
			b.WriteString(e.FormatProto2(""))
			b.WriteString("\n\n")
			report.TotalEnums++
		}

		// Sort and write Messages
		sort.Slice(msgs, func(i, j int) bool {
			return msgs[i].Name < msgs[j].Name
		})
		for _, m := range msgs {
			b.WriteString(m.FormatProto2(""))
			b.WriteString("\n\n")
			report.TotalMessages++
		}

		if err := os.WriteFile(targetFilePath, []byte(strings.TrimRight(b.String(), "\n")+"\n"), 0644); err != nil {
			return nil, fmt.Errorf("failed writing %s: %w", targetFilePath, err)
		}

		report.GeneratedFiles = append(report.GeneratedFiles, targetFilePath)
	}

	return report, nil
}

func collectMessageTypes(m *ast.MessageDef, acc map[string]bool) {
	for _, f := range m.Fields {
		cleanType := strings.TrimPrefix(f.Type, ".")
		if !isScalarType(cleanType) {
			acc[cleanType] = true
		}
		if f.IsMap {
			if !isScalarType(f.MapKey) {
				acc[f.MapKey] = true
			}
			if !isScalarType(f.MapValue) {
				acc[f.MapValue] = true
			}
		}
	}
	for _, o := range m.Oneofs {
		for _, f := range o.Fields {
			cleanType := strings.TrimPrefix(f.Type, ".")
			if !isScalarType(cleanType) {
				acc[cleanType] = true
			}
		}
	}
	for _, nm := range m.NestedMessages {
		collectMessageTypes(nm, acc)
	}
}

func isScalarType(t string) bool {
	switch t {
	case "double", "float", "int32", "int64", "uint32", "uint64",
		"sint32", "sint64", "fixed32", "fixed64", "sfixed32", "sfixed64",
		"bool", "string", "bytes":
		return true
	default:
		return false
	}
}

// UpdateClientPayloadVersion updates the waVersion variable in clientpayload.go.
func UpdateClientPayloadVersion(filePath string, fullVersion string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	parts := strings.Split(fullVersion, ".")
	if len(parts) < 3 {
		return fmt.Errorf("invalid version string format: %s (expected e.g. 2.3000.1046738589)", fullVersion)
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	build, _ := strconv.Atoi(parts[2])

	re := regexp.MustCompile(`var waVersion = WAVersionContainer\{\d+,\s*\d+,\s*\d+\}`)
	newDeclaration := fmt.Sprintf("var waVersion = WAVersionContainer{%d, %d, %d}", major, minor, build)

	if !re.Match(content) {
		return fmt.Errorf("waVersion declaration not found in %s", filePath)
	}

	updated := re.ReplaceAll(content, []byte(newDeclaration))
	return os.WriteFile(filePath, updated, 0644)
}
