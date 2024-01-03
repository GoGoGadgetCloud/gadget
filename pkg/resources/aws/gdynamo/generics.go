package gdynamo

import (
	"fmt"
	"reflect"
)

type (
	attributeType string
	attribute     struct {
		name          string
		attributeType attributeType
	}
	fieldPredicate func(string, reflect.Type) bool
)

var (
	attributeTypeString attributeType = "S"
	attributeTypeNumber attributeType = "N"
	attributeTypeBinary attributeType = "B"

	allowAll fieldPredicate = func(string, reflect.Type) bool {
		return true
	}
)

func skipKnownFields(t reflect.Type) fieldPredicate {
	blacklist := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		blacklist[i] = t.Field(i).Name
	}

	return func(name string, fieldType reflect.Type) bool {
		for _, field := range blacklist {
			if field == name {
				return true
			}
		}
		return false
	}
}

func getPrimaryKeyAttributeDefinitions(t reflect.Type) ([]attribute, error) {
	return mapStructToFields(t, allowAll)
}

func geTableAttributeDefinitions(t reflect.Type, id reflect.Type) ([]attribute, error) {
	return mapStructToFields(t, skipKnownFields(id))
}

func mapStructToFields(t reflect.Type, filter fieldPredicate) ([]attribute, error) {
	attributes := make([]attribute, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		nameTag := field.Tag.Get("attribute")
		if nameTag != "" {
			name = nameTag
		}
		if !filter(name, field.Type) {
			attributeType, err := mapFieldToAttributeType(field.Type)
			if err != nil {
				return nil, err
			}
			attributes = append(attributes, attribute{
				name:          name,
				attributeType: *attributeType,
			})
		}
	}
	return attributes, nil
}

func mapFieldToAttributeType(fieldType reflect.Type) (*attributeType, error) {
	switch fieldType.Kind() {
	case reflect.String:
		return &attributeTypeString, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &attributeTypeNumber, nil
	case reflect.Float32, reflect.Float64:
		return &attributeTypeNumber, nil
	case reflect.Struct:
		return &attributeTypeBinary, nil
	case reflect.Slice:
		if fieldType.Elem().Kind() == reflect.Uint8 {
			return &attributeTypeBinary, nil
		} else {
			return nil, fmt.Errorf("unsupported slice type %s", fieldType)
		}
	default:
		return nil, fmt.Errorf("type %s is not supported", fieldType.Kind())
	}
}
