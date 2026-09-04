package extractor

import (
	"fmt"
	"strings"

	"github.com/dop251/goja/ast"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja/token"
	protoAst "github.com/Thruqe/wa-proto/pkg/ast"
)

// ModuleSchema represents the extracted protobuf definitions of a single module.
type ModuleSchema struct {
	Name      string
	Version   string
	Messages  []*protoAst.MessageDef
	Enums     []*protoAst.EnumDef
	Imports   []string
	CrossRefs []string
}

// Extractor extracts protobuf schemas from WhatsApp Web JavaScript bundle sources.
type Extractor struct {
	Version string
	Modules map[string]*ModuleSchema
}

// NewExtractor creates a new extractor instance.
func NewExtractor(version string) *Extractor {
	return &Extractor{
		Version: version,
		Modules: make(map[string]*ModuleSchema),
	}
}

// ProcessSource parses a JavaScript bundle source and extracts all protobuf modules.
func (e *Extractor) ProcessSource(src string) error {
	// Patch known Webpack split idiosyncrasies
	patched := strings.ReplaceAll(src, "LimitSharing$Trigger", "LimitSharing$TriggerType")

	prog, err := parser.ParseFile(nil, "", patched, 0)
	if err != nil {
		return fmt.Errorf("parsing JS source: %w", err)
	}

	for _, stmt := range prog.Body {
		e.inspectStatement(stmt)
	}

	return nil
}

func (e *Extractor) inspectStatement(stmt ast.Statement) {
	callExpr, ok := getCallExpr(stmt)
	if !ok || len(callExpr.ArgumentList) < 3 {
		return
	}

	modNameLit, ok := callExpr.ArgumentList[0].(*ast.StringLiteral)
	if !ok {
		return
	}
	moduleName := modNameLit.Value.String()

	fn, ok := callExpr.ArgumentList[2].(*ast.FunctionLiteral)
	if !ok || fn.Body == nil {
		return
	}

	// First pass: collect local enums (both var and let/const declarations)
	enums := make(map[string]*protoAst.EnumDef)
	for _, s := range fn.Body.List {
		switch node := s.(type) {
		case *ast.VariableStatement:
			for _, decl := range node.List {
				if ident, ok := decl.Target.(*ast.Identifier); ok {
					if obj, ok := decl.Initializer.(*ast.ObjectLiteral); ok {
						if enumDef := parseEnumObject(ident.Name.String(), obj); enumDef != nil {
							enums[ident.Name.String()] = enumDef
						}
					}
				}
			}
		case *ast.LexicalDeclaration:
			for _, decl := range node.List {
				if ident, ok := decl.Target.(*ast.Identifier); ok {
					if obj, ok := decl.Initializer.(*ast.ObjectLiteral); ok {
						if enumDef := parseEnumObject(ident.Name.String(), obj); enumDef != nil {
							enums[ident.Name.String()] = enumDef
						}
					}
				}
			}
		}
	}

	// Second pass: collect message specs and module assignments
	var messages []*protoAst.MessageDef
	var exportedEnums []*protoAst.EnumDef

	for _, s := range fn.Body.List {
		switch node := s.(type) {
		case *ast.ExpressionStatement:
			if assign, ok := node.Expression.(*ast.AssignExpression); ok {
				if mem, ok := assign.Left.(*ast.DotExpression); ok {
					propName := mem.Identifier.Name.String()

					// Check if this exports an enum: f.EnumName = EnumName
					if rIdent, ok := assign.Right.(*ast.Identifier); ok {
						if enumDef, ok := enums[rIdent.Name.String()]; ok {
							enumDef.Name = propName
							exportedEnums = append(exportedEnums, enumDef)
							continue
						}
					}

					// Check if this is an inline enum assignment: f.EnumName = { ... }
					if obj, ok := assign.Right.(*ast.ObjectLiteral); ok && !strings.HasSuffix(propName, "Spec") {
						if enumDef := parseEnumObject(propName, obj); enumDef != nil {
							exportedEnums = append(exportedEnums, enumDef)
							continue
						}
					}

					// Check if this is a message spec: f.MessageNameSpec = { ... }
					if strings.HasSuffix(propName, "Spec") {
						msgName := strings.TrimSuffix(propName, "Spec")
						if obj, ok := assign.Right.(*ast.ObjectLiteral); ok {
							msgDef := parseMessageSpecObject(msgName, obj, enums)
							if msgDef != nil {
								messages = append(messages, msgDef)
							}
						}
					}
				}
			}
		}
	}

	if len(messages) > 0 || len(exportedEnums) > 0 {
		mod := e.Modules[moduleName]
		if mod == nil {
			mod = &ModuleSchema{
				Name:    moduleName,
				Version: e.Version,
			}
			e.Modules[moduleName] = mod
		}
		mod.Messages = append(mod.Messages, messages...)
		mod.Enums = append(mod.Enums, exportedEnums...)
	}
}

func parseEnumObject(name string, obj *ast.ObjectLiteral) *protoAst.EnumDef {
	var values []*protoAst.EnumValueDef
	for _, prop := range obj.Value {
		if pk, ok := prop.(*ast.PropertyKeyed); ok {
			key := getPropKey(pk.Key)
			if num, ok := getNumber(pk.Value); ok {
				values = append(values, &protoAst.EnumValueDef{
					Name: key,
					ID:   num,
				})
			} else {
				// Not a numeric enum
				return nil
			}
		}
	}
	if len(values) == 0 {
		return nil
	}
	return &protoAst.EnumDef{
		Name:   name,
		Values: values,
	}
}

func parseMessageSpecObject(msgName string, obj *ast.ObjectLiteral, localEnums map[string]*protoAst.EnumDef) *protoAst.MessageDef {
	msgDef := &protoAst.MessageDef{
		Name: msgName,
	}

	oneofsMap := make(map[string][]string)

	for _, prop := range obj.Value {
		pk, ok := prop.(*ast.PropertyKeyed)
		if !ok {
			continue
		}
		key := getPropKey(pk.Key)

		// Parse __oneofs__ constraint
		if key == "__oneofs__" {
			if oObj, ok := pk.Value.(*ast.ObjectLiteral); ok {
				for _, op := range oObj.Value {
					if opk, ok := op.(*ast.PropertyKeyed); ok {
						oneofName := getPropKey(opk.Key)
						if arr, ok := opk.Value.(*ast.ArrayLiteral); ok {
							for _, el := range arr.Value {
								if sLit, ok := el.(*ast.StringLiteral); ok {
									oneofsMap[oneofName] = append(oneofsMap[oneofName], sLit.Value.String())
								}
							}
						}
					}
				}
			}
			continue
		}

		if strings.HasPrefix(key, "__") {
			continue
		}

		arr, ok := pk.Value.(*ast.ArrayLiteral)
		if !ok || len(arr.Value) < 2 {
			continue
		}

		id, ok := getNumber(arr.Value[0])
		if !ok {
			continue
		}

		fieldType, rule, packed, isMap, mapKey, mapVal := parseTypeAndFlags(arr.Value[1], arr.Value, localEnums)

		msgDef.Fields = append(msgDef.Fields, &protoAst.FieldDef{
			Name:     key,
			ID:       id,
			Type:     fieldType,
			Rule:     rule,
			Packed:   packed,
			IsMap:    isMap,
			MapKey:   mapKey,
			MapValue: mapVal,
		})
	}

	// Assemble oneofs
	if len(oneofsMap) > 0 {
		fieldMap := make(map[string]*protoAst.FieldDef)
		for _, f := range msgDef.Fields {
			fieldMap[f.Name] = f
		}

		var remainingFields []*protoAst.FieldDef
		inOneof := make(map[string]bool)

		for oName, fNames := range oneofsMap {
			oDef := &protoAst.OneofDef{Name: oName}
			for _, fn := range fNames {
				if f, ok := fieldMap[fn]; ok {
					oDef.Fields = append(oDef.Fields, f)
					inOneof[fn] = true
				}
			}
			msgDef.Oneofs = append(msgDef.Oneofs, oDef)
		}

		for _, f := range msgDef.Fields {
			if !inOneof[f.Name] {
				remainingFields = append(remainingFields, f)
			}
		}
		msgDef.Fields = remainingFields
	}

	return msgDef
}

func parseTypeAndFlags(typeExpr ast.Expression, allElements []ast.Expression, localEnums map[string]*protoAst.EnumDef) (fieldType, rule string, packed, isMap bool, mapKey, mapVal string) {
	rule = "optional"
	var parts []ast.Expression
	unwrapBinaryOr(typeExpr, &parts)

	var detectedType string
	for _, p := range parts {
		if dot, ok := p.(*ast.DotExpression); ok {
			memName := dot.Identifier.Name.String()
			if memObj, ok := dot.Left.(*ast.DotExpression); ok {
				category := memObj.Identifier.Name.String()
				if category == "FLAGS" {
					switch strings.ToUpper(memName) {
					case "REPEATED":
						rule = "repeated"
					case "PACKED":
						packed = true
					}
				} else if category == "TYPES" {
					detectedType = strings.ToLower(memName)
				}
			} else {
				if strings.Contains(strings.ToUpper(memName), "PACKED") {
					packed = true
				}
				if strings.Contains(strings.ToUpper(memName), "REPEATED") {
					rule = "repeated"
				}
			}
		}
	}

	switch detectedType {
	case "map":
		isMap = true
		if len(allElements) > 2 {
			if arr, ok := allElements[2].(*ast.ArrayLiteral); ok && len(arr.Value) >= 2 {
				mapKey = getTypeNameFromExpr(arr.Value[0])
				mapVal = getTypeNameFromExpr(arr.Value[1])
			}
		}
		fieldType = fmt.Sprintf("map<%s, %s>", mapKey, mapVal)
		return

	case "enum", "message":
		if len(allElements) > 2 {
			target := getTypeNameFromExpr(allElements[2])
			if target != "" {
				fieldType = target
				return
			}
		}
		fieldType = "bytes"
		return

	case "":
		fieldType = "string"
		return

	default:
		fieldType = detectedType
		return
	}
}

func getTypeNameFromExpr(expr ast.Expression) string {
	switch e := expr.(type) {
	case *ast.Identifier:
		name := e.Name.String()
		return strings.TrimSuffix(name, "Spec")
	case *ast.DotExpression:
		name := e.Identifier.Name.String()
		return strings.TrimSuffix(name, "Spec")
	case *ast.StringLiteral:
		return strings.TrimSuffix(e.Value.String(), "Spec")
	default:
		return ""
	}
}

func unwrapBinaryOr(expr ast.Expression, acc *[]ast.Expression) {
	if bin, ok := expr.(*ast.BinaryExpression); ok && bin.Operator == token.OR {
		unwrapBinaryOr(bin.Left, acc)
		unwrapBinaryOr(bin.Right, acc)
	} else if expr != nil {
		*acc = append(*acc, expr)
	}
}

func getCallExpr(stmt ast.Statement) (*ast.CallExpression, bool) {
	if exprStmt, ok := stmt.(*ast.ExpressionStatement); ok {
		if call, ok := exprStmt.Expression.(*ast.CallExpression); ok {
			return call, true
		}
	}
	return nil, false
}

func getNumber(expr ast.Expression) (int, bool) {
	switch e := expr.(type) {
	case *ast.NumberLiteral:
		switch v := e.Value.(type) {
		case int:
			return v, true
		case int64:
			return int(v), true
		case float64:
			return int(v), true
		}
	case *ast.UnaryExpression:
		if e.Operator == token.MINUS {
			if n, ok := getNumber(e.Operand); ok {
				return -n, true
			}
		}
	}
	return 0, false
}

func getPropKey(expr ast.Expression) string {
	switch e := expr.(type) {
	case *ast.Identifier:
		return e.Name.String()
	case *ast.StringLiteral:
		return e.Value.String()
	default:
		return ""
	}
}
