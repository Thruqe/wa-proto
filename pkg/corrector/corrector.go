package corrector

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Thruqe/wa-proto/pkg/ast"
)

var (
	requiredRe = regexp.MustCompile(`\brequired\s+`)
	enumRe     = regexp.MustCompile(`^enum\s+([A-Za-z0-9_]+)\s*\{`)
	enumValRe  = regexp.MustCompile(`^(\s*)([A-Za-z0-9_]+)\s*=\s*(-?\d+)\s*;(.*)$`)
)

// FixProto3 corrects a ProtoSchema so it conforms strictly to proto3 syntax.
func FixProto3(schema *ast.ProtoSchema) {
	// 1. Fix messages: replace any required rule with optional
	for _, m := range schema.Messages {
		fixMessageProto3(m)
	}

	// 2. Fix enums: ensure first value is 0 and handle duplicates
	usedEnumConstants := make(map[string]string) // constantName -> enumName
	for _, e := range schema.Enums {
		fixEnumProto3(e, usedEnumConstants)
	}
}

func fixMessageProto3(m *ast.MessageDef) {
	for _, f := range m.Fields {
		if f.Rule == "required" {
			f.Rule = "optional"
		}
	}
	for _, nm := range m.NestedMessages {
		fixMessageProto3(nm)
	}
	usedNestedEnumConstants := make(map[string]string)
	for _, ne := range m.NestedEnums {
		fixEnumProto3(ne, usedNestedEnumConstants)
	}
}

func fixEnumProto3(e *ast.EnumDef, usedConstants map[string]string) {
	if len(e.Values) == 0 {
		e.Values = append(e.Values, &ast.EnumValueDef{
			Name: fmt.Sprintf("%s_UNKNOWN", strings.ToUpper(e.Name)),
			ID:   0,
		})
		return
	}

	// Check if a zero value already exists
	zeroIdx := -1
	for idx, v := range e.Values {
		if v.ID == 0 {
			zeroIdx = idx
			break
		}
	}

	if zeroIdx == 0 {
		// Already has 0 as first element
	} else if zeroIdx > 0 {
		// Move the 0 value to the front
		zeroVal := e.Values[zeroIdx]
		e.Values = append([]*ast.EnumValueDef{zeroVal}, append(e.Values[:zeroIdx], e.Values[zeroIdx+1:]...)...)
	} else {
		// No 0 value exists, prepend an UNKNOWN/DEFAULT value = 0
		unknownName := fmt.Sprintf("%s_UNKNOWN", strings.ToUpper(e.Name))
		if _, exists := usedConstants[unknownName]; exists {
			unknownName = fmt.Sprintf("%s_UNSPECIFIED", strings.ToUpper(e.Name))
		}
		e.Values = append([]*ast.EnumValueDef{{
			Name: unknownName,
			ID:   0,
		}}, e.Values...)
	}

	// Resolve enum value collisions in the package scope
	for _, v := range e.Values {
		if prevEnum, exists := usedConstants[v.Name]; exists && prevEnum != e.Name {
			// Prefix constant with EnumName
			v.Name = fmt.Sprintf("%s_%s", strings.ToUpper(e.Name), v.Name)
		}
		usedConstants[v.Name] = e.Name
	}
}

// FixProto3Content fixes raw proto3 content strings directly.
func FixProto3Content(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	inEnum := false
	currentEnumName := ""
	firstValSeen := false
	hasZeroVal := false
	var enumLines []string
	usedEnumValues := make(map[string]string)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Fix 1: Replace required with optional
		if strings.Contains(line, "required ") {
			line = requiredRe.ReplaceAllString(line, "optional ")
			trimmed = strings.TrimSpace(line)
		}

		// Check for enum start
		if em := enumRe.FindStringSubmatch(trimmed); len(em) > 1 {
			inEnum = true
			currentEnumName = em[1]
			firstValSeen = false
			hasZeroVal = false
			enumLines = []string{line}
			continue
		}

		if inEnum {
			if trimmed == "}" || strings.HasPrefix(trimmed, "};") {
				// Process enum lines before flushing
				processed := processEnumLines(currentEnumName, enumLines, hasZeroVal, usedEnumValues)
				result = append(result, processed...)
				result = append(result, line)
				inEnum = false
				currentEnumName = ""
				continue
			}

			if vm := enumValRe.FindStringSubmatch(line); len(vm) > 3 {
				val, _ := strconv.Atoi(vm[3])
				if val == 0 {
					hasZeroVal = true
				}
				_ = firstValSeen
				firstValSeen = true
			}
			enumLines = append(enumLines, line)
			continue
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

func processEnumLines(enumName string, lines []string, hasZero bool, usedValues map[string]string) []string {
	if len(lines) <= 1 {
		return lines
	}

	header := lines[0]
	valLines := lines[1:]

	var fixedValLines []string
	zeroLineIdx := -1

	for idx, l := range valLines {
		if vm := enumValRe.FindStringSubmatch(l); len(vm) > 3 {
			val, _ := strconv.Atoi(vm[3])
			if val == 0 {
				zeroLineIdx = idx
			}
		}
	}

	if !hasZero {
		// Prepend a zero value
		indent := "    "
		unknownName := fmt.Sprintf("%s_UNKNOWN", strings.ToUpper(enumName))
		if _, exists := usedValues[unknownName]; exists {
			unknownName = fmt.Sprintf("%s_UNSPECIFIED", strings.ToUpper(enumName))
		}
		usedValues[unknownName] = enumName
		fixedValLines = append(fixedValLines, fmt.Sprintf("%s%s = 0;", indent, unknownName))
	} else if zeroLineIdx > 0 {
		// Move the zero line to index 0
		zeroLine := valLines[zeroLineIdx]
		valLines = append([]string{zeroLine}, append(valLines[:zeroLineIdx], valLines[zeroLineIdx+1:]...)...)
	}

	for _, l := range valLines {
		if vm := enumValRe.FindStringSubmatch(l); len(vm) > 4 {
			indent := vm[1]
			name := vm[2]
			val := vm[3]
			trailing := vm[4]

			// Disambiguate duplicate enum constants in same file
			if prevEnum, exists := usedValues[name]; exists && prevEnum != enumName {
				name = fmt.Sprintf("%s_%s", strings.ToUpper(enumName), name)
			}
			usedValues[name] = enumName
			fixedValLines = append(fixedValLines, fmt.Sprintf("%s%s = %s;%s", indent, name, val, trailing))
		} else {
			fixedValLines = append(fixedValLines, l)
		}
	}

	return append([]string{header}, fixedValLines...)
}
