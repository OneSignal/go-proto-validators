// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

/*

The validator plugin generates a Validate method for each message.
By default, if none of the message's fields are annotated with the gogo validator annotation, it returns a nil.
In case some of the fields are annotated, the Validate function returns nil upon sucessful validation, or an error
describing why the validation failed.
The Validate method is called recursively for all submessage of the message.

TODO(michal): ADD COMMENTS.

Equal is enabled using the following extensions:

  - equal
  - equal_all

While VerboseEqual is enable dusing the following extensions:

  - verbose_equal
  - verbose_equal_all

The equal plugin also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

Let us look at:

  github.com/gogo/protobuf/test/example/example.proto

Btw all the output can be seen at:

  github.com/gogo/protobuf/test/example/*

The following message:



given to the equal plugin, will generate the following code:



and the following test code:


*/
package plugin

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"

	validator "github.com/OneSignal/go-proto-validators"
)

const uuidPattern = "^([a-fA-F0-9]{8}-" +
	"[a-fA-F0-9]{4}-" +
	"[%s][a-fA-F0-9]{3}-" +
	"[8|9|aA|bB][a-fA-F0-9]{3}-" +
	"[a-fA-F0-9]{12})?$"

type plugin struct {
	*generator.Generator
	generator.PluginImports
	regexPkg      generator.Single
	fmtPkg        generator.Single
	validatorPkg  generator.Single
	uuidPkg       generator.Single
	useGogoImport bool
}

func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "validator"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	if !p.useGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.regexPkg = p.NewImport("regexp")
	p.fmtPkg = p.NewImport("fmt")
	p.validatorPkg = p.NewImport("github.com/OneSignal/go-proto-validators")
	p.uuidPkg = p.NewImport("github.com/gofrs/uuid")

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.generateRegexVars(file, msg)
		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		} else {
			p.generateProto2Message(file, msg)
		}
	}
}

func getFieldValidatorIfAny(field *descriptor.FieldDescriptorProto) *validator.FieldValidator {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, validator.E_Field)
		if err == nil && v.(*validator.FieldValidator) != nil {
			return (v.(*validator.FieldValidator))
		}
	}
	return nil
}

func getOneofValidatorIfAny(oneof *descriptor.OneofDescriptorProto) *validator.OneofValidator {
	if oneof.Options != nil {
		v, err := proto.GetExtension(oneof.Options, validator.E_Oneof)
		if err == nil && v.(*validator.OneofValidator) != nil {
			return (v.(*validator.OneofValidator))
		}
	}
	return nil
}

func (p *plugin) isSupportedInt(field *descriptor.FieldDescriptorProto) bool {
	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_INT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_UINT32, descriptor.FieldDescriptorProto_TYPE_UINT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_SINT32, descriptor.FieldDescriptorProto_TYPE_SINT64:
		return true
	}
	return false
}

func (p *plugin) isSupportedFloat(field *descriptor.FieldDescriptorProto) bool {
	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_FLOAT, descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return true
	case descriptor.FieldDescriptorProto_TYPE_FIXED32, descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32, descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return true
	}
	return false
}

func (p *plugin) generateRegexVars(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {
		v := getFieldValidatorIfAny(field)
		if v != nil {
			fieldName := p.GetOneOfFieldName(message, field)
			protocolName := *field.Name
			if v.Regex != nil && v.UuidVer != nil {
				fmt.Fprintf(os.Stderr, "WARNING: regex and uuid v is set for field %v.%v, only one of them can be set. Regex and UUID v is ignored for this field.", ccTypeName, protocolName)
			} else if v.UuidVer != nil {
				uuid, err := getUUIDRegex(&v.UuidVer.Value)
				if err != nil {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v error %s.\n", ccTypeName, protocolName, err)
				} else {
					v.Regex = &validator.StringValue{Value: uuid}
					p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "`", v.Regex.Value, "`", `)`)
				}
			} else if v.Regex != nil {
				p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "`", v.Regex.Value, "`", `)`)
			}
		}
	}
}

func (p *plugin) GetFieldName(message *generator.Descriptor, field *descriptor.FieldDescriptorProto) string {
	fieldName := p.Generator.GetFieldName(message, field)
	if p.useGogoImport {
		fieldName := p.Generator.GetFieldName(message, field)
		return fieldName
	}
	if gogoproto.IsEmbed(field) {
		fieldName = generator.CamelCase(*field.Name)
	}
	return fieldName
}

func (p *plugin) GetOneOfFieldName(message *generator.Descriptor, field *descriptor.FieldDescriptorProto) string {
	fieldName := p.Generator.GetOneOfFieldName(message, field)
	if p.useGogoImport {
		return fieldName
	}
	if gogoproto.IsEmbed(field) {
		fieldName = generator.CamelCase(*field.Name)
	}
	return fieldName
}

func (p *plugin) generateProto2Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())

	p.P(`func (this *`, ccTypeName, `) Validate() error {`)
	p.In()
	for _, field := range message.Field {
		fieldName := p.GetFieldName(message, field)
		fieldValidator := getFieldValidatorIfAny(field)
		if fieldValidator == nil && !field.IsMessage() {
			continue
		}
		protocolName := *field.Name
		if p.validatorWithMessageExists(fieldValidator) {
			fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is a proto2 message, validator.msg_exists has no effect\n", ccTypeName, protocolName)
		}
		variableName := "this." + fieldName
		repeated := field.IsRepeated()
		nullable := gogoproto.IsNullable(field) && !(p.useGogoImport && gogoproto.IsEmbed(field))
		// For proto2 syntax, only Gogo generates non-pointer fields
		nonpointer := gogoproto.ImportsGoGoProto(file.FileDescriptorProto) && !gogoproto.IsNullable(field)
		if repeated {
			p.generateRepeatedCountValidator(variableName, ccTypeName, fieldName, protocolName, fieldValidator)
			if field.IsMessage() || p.validatorWithNonRepeatedConstraint(fieldValidator) {
				p.P(`for index, item := range `, variableName, `{`)
				p.In()
				variableName = "item"
			}
		} else if nullable {
			p.P(`if `, variableName, ` != nil {`)
			p.In()
			if !field.IsBytes() {
				variableName = "*(" + variableName + ")"
			}
		} else if nonpointer {
			// can use the field directly
		} else if !field.IsMessage() {
			variableName = `this.Get` + fieldName + `()`
		}
		if !repeated && fieldValidator != nil {
			if fieldValidator.RepeatedCountMin != nil {
				fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is not repeated, validator.min_elts has no effects\n", ccTypeName, fieldName)
			}
			if fieldValidator.RepeatedCountMax != nil {
				fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is not repeated, validator.max_elts has no effects\n", ccTypeName, fieldName)
			}
		}
		if field.IsString() {
			p.generateStringValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if p.isSupportedInt(field) {
			p.generateIntValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if field.IsEnum() {
			p.generateEnumValidator(field, variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if p.isSupportedFloat(field) {
			p.generateFloatValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if field.IsBytes() {
			p.generateLengthValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if field.IsMessage() {
			if repeated && nullable {
				variableName = "*(item)"
			}
			p.P(`if err := `, p.validatorPkg.Use(), `.CallValidatorIfExists(&(`, variableName, `)); err != nil {`)
			p.In()
			p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `", err)`)
			p.Out()
			p.P(`}`)
		}
		if repeated {
			// end the repeated loop
			if field.IsMessage() || p.validatorWithNonRepeatedConstraint(fieldValidator) {
				// generate some dummy code to trick the compiler in case we don't
				// use index as part of an error message.
				p.P(`// index += 1 is some dummy code to trick the compiler in case we don't use`)
				p.P(`// idex as part of an error message`)
				p.P(`index += 1`)
				// This internal 'if' cannot be refactored as it would change semantics with respect to the corresponding prelude 'if's
				p.Out()
				p.P(`}`)
			}
		} else if nullable {
			// end the if around nullable
			p.Out()
			p.P(`}`)
		}
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (this *`, ccTypeName, `) Validate() error {`)
	p.In()

	for _, oneof := range message.OneofDecl {
		oneofValidator := getOneofValidatorIfAny(oneof)
		if oneofValidator == nil {
			continue
		}
		if oneofValidator.GetRequired() {
			oneOfName := generator.CamelCase(oneof.GetName())
			p.P(`if this.Get` + oneOfName + `() == nil {`)
			p.In()
			p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, oneof.GetName(), `",`, p.fmtPkg.Use(), `.Errorf("one of the fields must be set"))`)
			p.Out()
			p.P(`}`)
		}
	}
	for _, field := range message.Field {
		fieldValidator := getFieldValidatorIfAny(field)
		if fieldValidator == nil && !field.IsMessage() {
			continue
		}
		isOneOf := field.OneofIndex != nil
		fieldName := p.GetOneOfFieldName(message, field)
		protocolName := *field.Name
		variableName := "this." + fieldName
		repeated := field.IsRepeated()
		// Golang's proto3 has no concept of unset primitive fields
		nullable := (gogoproto.IsNullable(field) || !gogoproto.ImportsGoGoProto(file.FileDescriptorProto)) && field.IsMessage() && !(p.useGogoImport && gogoproto.IsEmbed(field))
		if p.fieldIsProto3Map(file, message, field) {
			p.P(`// Validation of proto3 map<> fields is unsupported.`)
			continue
		}
		if isOneOf {
			p.In()
			oneOfName := p.GetFieldName(message, field)
			oneOfType := p.OneOfTypeName(message, field)
			// if x, ok := m.GetType().(*OneOfMessage3_OneInt); ok {
			p.P(`if oneOfNester, ok := this.Get` + oneOfName + `().(* ` + oneOfType + `); ok {`)
			variableName = "oneOfNester." + p.GetOneOfFieldName(message, field)
		}
		if repeated {
			p.generateRepeatedCountValidator(variableName, ccTypeName, fieldName, protocolName, fieldValidator)
			if field.IsMessage() || p.validatorWithNonRepeatedConstraint(fieldValidator) {
				p.P(`for index, item := range `, variableName, `{`)
				p.In()
				variableName = "item"
			}
		} else if fieldValidator != nil {
			if fieldValidator.RepeatedCountMin != nil {
				fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is not repeated, validator.min_elts has no effects\n", ccTypeName, protocolName)
			}
			if fieldValidator.RepeatedCountMax != nil {
				fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is not repeated, validator.max_elts has no effects\n", ccTypeName, protocolName)
			}
		}
		if field.IsString() {
			p.generateStringValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if p.isSupportedInt(field) {
			p.generateIntValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if field.IsEnum() {
			p.generateEnumValidator(field, variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if p.isSupportedFloat(field) {
			p.generateFloatValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if field.IsBytes() {
			p.generateLengthValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fieldValidator)
		} else if field.IsMessage() {
			if p.validatorWithOnesignalUuidV1(fieldValidator) {
				if field.TypeName != nil && *field.TypeName == ".uuid.v1.UUID" {
					p.P(`if nil == `, variableName, `{`)
					p.In()
					if repeated {
						p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `",`, p.fmtPkg.Use(), `.Errorf("nil UUID at index %d", index))`)
					} else {
						p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `",`, p.fmtPkg.Use(), `.Errorf("nil UUID"))`)
					}
					p.Out()
					p.P(`} else {`)
					p.In()
					p.P(`if _, err := `, p.uuidPkg.Use(), `.FromBytes(`, variableName, `.Data); err != nil {`)
					p.In()
					if repeated {
						p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `",`, p.fmtPkg.Use(), `.Errorf("%v at index %d", err, index))`)
					} else {
						p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `", err)`)
					}
					p.Out()
					p.P(`}`)
					p.Out()
					p.P(`}`)
				} else {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is not 'uuid.v1.UUID', validator.onesignal_uuid_v1 has no effect\n", ccTypeName, fieldName)
				}
			}
			if p.validatorWithMessageExists(fieldValidator) {
				if nullable && !repeated {
					p.P(`if nil == `, variableName, `{`)
					p.In()
					p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `",`, p.fmtPkg.Use(), `.Errorf("message must exist"))`)
					p.Out()
					p.P(`}`)
				} else if repeated {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is repeated, validator.msg_exists has no effect\n", ccTypeName, fieldName)
				} else if !nullable {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is a nullable=false, validator.msg_exists has no effect\n", ccTypeName, fieldName)
				}
			}
			if nullable {
				p.P(`if `, variableName, ` != nil {`)
				p.In()
			} else {
				// non-nullable fields in proto3 store actual structs, we need pointers to operate on interfaces
				variableName = "&(" + variableName + ")"
			}
			p.P(`if err := `, p.validatorPkg.Use(), `.CallValidatorIfExists(`, variableName, `); err != nil {`)
			p.In()
			p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `", err)`)
			p.Out()
			p.P(`}`)
			if nullable {
				p.Out()
				p.P(`}`)
			}
		}
		if repeated && (field.IsMessage() || p.validatorWithNonRepeatedConstraint(fieldValidator)) {
			// generate some dummy code to trick the compiler in case we don't
			// use index as part of an error message.
			p.P(`// index += 1 is some dummy code to trick the compiler in case we don't use`)
			p.P(`// idex as part of an error message`)
			p.P(`index += 1`)
			// end the repeated loop
			p.Out()
			p.P(`}`)
		}
		if isOneOf {
			// end the oneof if statement
			p.Out()
			p.P(`}`)
		}
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
}

func (p *plugin) generateIntValidator(variableName, ccTypeName, fieldName, protocolName string, repeated bool, fv *validator.FieldValidator) {
	if fv.IntGt != nil {
		p.P(`if !(`, variableName, ` > `, &(fv.IntGt.Value), `) {`)
		p.In()
		errorStr := fmt.Sprintf(`be greater than '%d'`, fv.GetIntGt().GetValue())
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	if fv.IntLt != nil {
		p.P(`if !(`, variableName, ` < `, &(fv.IntLt.Value), `) {`)
		p.In()
		errorStr := fmt.Sprintf(`be less than '%d'`, fv.GetIntLt().GetValue())
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateEnumValidator(
	field *descriptor.FieldDescriptorProto,
	variableName, ccTypeName, fieldName, protocolName string,
	repeated bool,
	fv *validator.FieldValidator,
) {
	if fv.GetIsInEnum() {
		enum := p.ObjectNamed(field.GetTypeName()).(*generator.EnumDescriptor)
		p.P(`if _, ok := `, strings.Join(enum.TypeName(), "_"), "_name[int32(", variableName, ")]; !ok {")
		p.In()
		p.generateErrorString(variableName, protocolName, repeated, fmt.Sprintf("be a valid %s field", strings.Join(enum.TypeName(), "_")), fv)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateLengthValidator(variableName, ccTypeName, fieldName, protocolName string, repeated bool, fv *validator.FieldValidator) {
	if fv.LengthGt != nil {
		p.P(`if !( len(`, variableName, `) > `, &(fv.LengthGt.Value), `) {`)
		p.In()
		errorStr := fmt.Sprintf(`have a length greater than '%d'`, fv.GetLengthGt().GetValue())
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}

	if fv.LengthLt != nil {
		p.P(`if !( len(`, variableName, `) < `, &(fv.LengthLt.Value), `) {`)
		p.In()
		errorStr := fmt.Sprintf(`have a length smaller than '%d'`, fv.GetLengthLt().GetValue())
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}

	if fv.LengthEq != nil {
		p.P(`if !( len(`, variableName, `) == `, &(fv.LengthEq.Value), `) {`)
		p.In()
		errorStr := fmt.Sprintf(`have a length equal to '%d'`, fv.GetLengthEq().GetValue())
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateFloatValidator(variableName, ccTypeName, fieldName, protocolName string, repeated bool, fv *validator.FieldValidator) {
	upperIsStrict := true
	lowerIsStrict := true

	// First check for incompatible constraints (i.e flt_lt & flt_lte both defined, etc) and determine the real limits.
	if fv.FloatEpsilon != nil && fv.FloatLt == nil && fv.FloatGt == nil {
		fmt.Fprintf(os.Stderr, "WARNING: field %v.%v has no 'float_lt' or 'float_gt' field so setting 'float_epsilon' has no effect.", ccTypeName, fieldName)
	}
	if fv.FloatLt != nil && fv.FloatLte != nil {
		fmt.Fprintf(os.Stderr, "WARNING: field %v.%v has both 'float_lt' and 'float_lte' constraints, only the strictest will be used.", ccTypeName, fieldName)
		strictLimit := fv.GetFloatLt().GetValue()
		if fv.FloatEpsilon != nil {
			strictLimit += fv.GetFloatEpsilon().GetValue()
		}

		if fv.GetFloatLte().GetValue() < strictLimit {
			upperIsStrict = false
		}
	} else if fv.FloatLte != nil {
		upperIsStrict = false
	}

	if fv.FloatGt != nil && fv.FloatGte != nil {
		fmt.Fprintf(os.Stderr, "WARNING: field %v.%v has both 'float_gt' and 'float_gte' constraints, only the strictest will be used.", ccTypeName, fieldName)
		strictLimit := fv.GetFloatGt().GetValue()
		if fv.FloatEpsilon != nil {
			strictLimit -= fv.GetFloatEpsilon().GetValue()
		}

		if fv.GetFloatGte().GetValue() > strictLimit {
			lowerIsStrict = false
		}
	} else if fv.FloatGte != nil {
		lowerIsStrict = false
	}

	// Generate the constraint checking code.
	errorStr := ""
	compareStr := ""
	if fv.FloatGt != nil || fv.FloatGte != nil {
		compareStr = fmt.Sprint(`if !(`, variableName)
		if lowerIsStrict {
			errorStr = fmt.Sprintf(`be strictly greater than '%g'`, fv.GetFloatGt().GetValue())
			if fv.FloatEpsilon != nil {
				errorStr += fmt.Sprintf(` with a tolerance of '%g'`, fv.GetFloatEpsilon().GetValue())
				compareStr += fmt.Sprint(` + `, fv.GetFloatEpsilon().GetValue())
			}
			compareStr += fmt.Sprint(` > `, fv.GetFloatGt().GetValue(), `) {`)
		} else {
			errorStr = fmt.Sprintf(`be greater than or equal to '%g'`, fv.GetFloatGte().GetValue())
			compareStr += fmt.Sprint(` >= `, fv.GetFloatGte().GetValue(), `) {`)
		}
		p.P(compareStr)
		p.In()
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}

	if fv.FloatLt != nil || fv.FloatLte != nil {
		compareStr = fmt.Sprint(`if !(`, variableName)
		if upperIsStrict {
			errorStr = fmt.Sprintf(`be strictly lower than '%g'`, fv.GetFloatLt().GetValue())
			if fv.FloatEpsilon != nil {
				errorStr += fmt.Sprintf(` with a tolerance of '%g'`, fv.GetFloatEpsilon().GetValue())
				compareStr += fmt.Sprint(` - `, fv.GetFloatEpsilon().GetValue())
			}
			compareStr += fmt.Sprint(` < `, fv.GetFloatLt().GetValue(), `) {`)
		} else {
			errorStr = fmt.Sprintf(`be lower than or equal to '%g'`, fv.GetFloatLte().GetValue())
			compareStr += fmt.Sprint(` <= `, fv.GetFloatLte().GetValue(), `) {`)
		}
		p.P(compareStr)
		p.In()
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
}

// getUUIDRegex returns a regex to validate that a string is in UUID
// format. The version parameter specified the UUID version. If version is 0,
// the returned regex is valid for any UUID version
func getUUIDRegex(version *int32) (string, error) {
	if version == nil {
		return "", nil
	} else if *version < 0 || *version > 5 {
		return "", fmt.Errorf("UUID version should be between 0-5, Got %d", *version)
	} else if *version == 0 {
		return fmt.Sprintf(uuidPattern, "1-5"), nil
	} else {
		return fmt.Sprintf(uuidPattern, strconv.Itoa(int(*version))), nil
	}
}

func (p *plugin) generateStringValidator(variableName, ccTypeName, fieldName, protocolName string, repeated bool, fv *validator.FieldValidator) {
	if fv.Regex != nil || fv.UuidVer != nil {
		if fv.UuidVer != nil {
			uuid, err := getUUIDRegex(&fv.UuidVer.Value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "WARNING: field %v.%v error %s.\n", ccTypeName, protocolName, err)
			} else {
				fv.Regex = &validator.StringValue{Value: uuid}
			}
		}
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
		p.In()
		errorStr := "be a string conforming to regex " + strconv.Quote(fv.GetRegex().GetValue())
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	if fv.GetStringNotEmpty() {
		p.P(`if `, variableName, ` == "" {`)
		p.In()
		errorStr := "not be an empty string"
		p.generateErrorString(variableName, protocolName, repeated, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	p.generateLengthValidator(variableName, ccTypeName, fieldName, protocolName, repeated, fv)
}

func (p *plugin) generateRepeatedCountValidator(variableName, ccTypeName, fieldName, protocolName string, fv *validator.FieldValidator) {
	if fv == nil {
		return
	}
	if fv.RepeatedCountMin != nil {
		compareStr := fmt.Sprint(`if len(`, variableName, `) < `, fv.GetRepeatedCountMin().GetValue(), ` {`)
		p.P(compareStr)
		p.In()
		errorStr := fmt.Sprint(`contain at least `, fv.GetRepeatedCountMin().GetValue(), ` elements`)
		p.generateErrorString(variableName, protocolName, false, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	if fv.RepeatedCountMax != nil {
		compareStr := fmt.Sprint(`if len(`, variableName, `) > `, fv.GetRepeatedCountMax().GetValue(), ` {`)
		p.P(compareStr)
		p.In()
		errorStr := fmt.Sprint(`contain at most `, fv.GetRepeatedCountMax().GetValue(), ` elements`)
		p.generateErrorString(variableName, protocolName, false, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateErrorString(variableName string, protocolName string, repeated bool, specificError string, fv *validator.FieldValidator) {
	if fv.GetHumanError() == "" {
		if repeated {
			p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `",`, p.fmtPkg.Use(), ".Errorf(`value '%v' must ", specificError, " at index %d`", `, `, variableName, `, index))`)
		} else {
			p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `",`, p.fmtPkg.Use(), ".Errorf(`value '%v' must ", specificError, "`", `, `, variableName, `))`)
		}
	} else {
		p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, protocolName, `",`, p.fmtPkg.Use(), ".Errorf(`", fv.GetHumanError(), "`))")
	}
}

func (p *plugin) fieldIsProto3Map(file *generator.FileDescriptor, message *generator.Descriptor, field *descriptor.FieldDescriptorProto) bool {
	// Context from descriptor.proto
	// Whether the message is an automatically generated map entry type for the
	// maps field.
	//
	// For maps fields:
	//     map<KeyType, ValueType> map_field = 1;
	// The parsed descriptor looks like:
	//     message MapFieldEntry {
	//         option map_entry = true;
	//         optional KeyType key = 1;
	//         optional ValueType value = 2;
	//     }
	//     repeated MapFieldEntry map_field = 1;
	//
	// Implementations may choose not to generate the map_entry=true message, but
	// use a native map in the target language to hold the keys and values.
	// The reflection APIs in such implementions still need to work as
	// if the field is a repeated message field.
	//
	// NOTE: Do not set the option in .proto files. Always use the maps syntax
	// instead. The option should only be implicitly set by the proto compiler
	// parser.
	if field.GetType() != descriptor.FieldDescriptorProto_TYPE_MESSAGE || !field.IsRepeated() {
		return false
	}
	typeName := field.GetTypeName()
	var msg *descriptor.DescriptorProto
	if strings.HasPrefix(typeName, ".") {
		// Fully qualified case, look up in global map, must work or fail badly.
		msg = p.ObjectNamed(field.GetTypeName()).(*generator.Descriptor).DescriptorProto
	} else {
		// Nested, relative case.
		msg = file.GetNestedMessage(message.DescriptorProto, field.GetTypeName())
	}
	return msg.GetOptions().GetMapEntry()
}

func (p *plugin) validatorWithMessageExists(fv *validator.FieldValidator) bool {
	return fv.GetMsgExists()
}

func (p *plugin) validatorWithOnesignalUuidV1(fv *validator.FieldValidator) bool {
	return fv.GetOnesignalUuidV1()
}

func (p *plugin) validatorWithNonRepeatedConstraint(fv *validator.FieldValidator) bool {
	if fv == nil {
		return false
	}

	// Need to use reflection in order to be future-proof for new types of constraints.
	v := reflect.ValueOf(*fv)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name

		// Identify non-repeated constraints based on their name.
		switch fieldName {
		case "RepeatedCountMin", "RepeatedCountMax":
			continue
		case "StringNotEmpty":
			// We only want to generate a for loop over the elements of the
			// repeated string field if the validator is enabled. Proto3 is
			// different in that primitive types are no longer optional, they
			// will have a zero value if not set.
			if v.Field(i).Bool() {
				return true
			} else {
				continue
			}
		default:
			// Most known validators will have a pointer type and we should skip any fields
			// that are not pointers (i.e unknown fields, etc) as well as 'nil' pointers that
			// don't lead to anything.
			if v.Type().Field(i).Type.Kind() != reflect.Ptr || v.Field(i).IsNil() {
				continue
			}
			return true
		}

	}
	p.P("// default false")
	return false
}

func (p *plugin) regexName(ccTypeName string, fieldName string) string {
	return "_regex_" + ccTypeName + "_" + fieldName
}
