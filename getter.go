package bvmgo_reflect

import (
	"fmt"
	"reflect"
	"strings"
)

// getFieldValue function returns structure "fieldName" field value.
//
// getFieldValue function returns an error if:
//   - sourceStructure is not a structure (or a map),
//   - sourceStructure is nil or invalid,
//   - fieldName is not a valid field name,
//   - fieldName is not found on sourceStructure structure (or map).
func getFieldValue(sourceStructure any, fieldName string) (fieldValue reflect.Value, err error) {
	cleanFieldName := strings.TrimSpace(fieldName)
	sourceValue := reflect.ValueOf(sourceStructure)
	// Check field name
	if len(strings.TrimSpace(cleanFieldName)) == 0 {
		err = fmt.Errorf("field name is empty")
		return
	}
	// Check is not null
	if !sourceValue.IsValid() ||
		((sourceValue.Kind() == reflect.Ptr || sourceValue.Kind() == reflect.Map) && sourceValue.IsNil()) {
		err = fmt.Errorf("a not nil pointer is required to get value from [%s] field",
			cleanFieldName)
		return
	}
	// If pointer, get pointed element
	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}
	// Check is a structure or a map
	if sourceValue.Kind() == reflect.Struct {
		fieldValue = sourceValue.FieldByName(cleanFieldName)
		// Check field exists
		if !fieldValue.IsValid() {
			err = fmt.Errorf("[%s.%s] field is not found", typeName(sourceValue.Type()), cleanFieldName)
			return
		}
		return
	} else if sourceValue.Kind() == reflect.Map {
		fieldValue = sourceValue.MapIndex(reflect.ValueOf(cleanFieldName))
		// Check map entry exists
		if !fieldValue.IsValid() {
			err = fmt.Errorf("[%s] map entry is not found", cleanFieldName)
			return
		}
		return
	} else {
		err = fmt.Errorf("unsupported type [%s], a structure or a map is required to get value from [%s] field",
			typeName(sourceValue.Type()), cleanFieldName)
		return
	}
}

// GetFieldString function returns structure "fieldName" field value.
//
// GetFieldString function returns an error if:
//   - sourceStructure is not a structure (or a map),
//   - sourceStructure is nil or invalid,
//   - fieldName is not a valid field name,
//   - fieldName is not found on sourceStructure structure (or map),
//   - fieldName is not a string.
func GetFieldString(sourceStructure any, fieldName string) (value string, err error) {
	fieldValue, err := getFieldValue(sourceStructure, fieldName)
	// Check field exists
	if err != nil {
		return
	}
	if fieldValue.Kind() == reflect.String {
		value = fieldValue.String()
		return
	} else {
		err = fmt.Errorf("[%s] field type [%s] is not a string",
			typeName(fieldValue.Type()), fieldName)
		return
	}
}
