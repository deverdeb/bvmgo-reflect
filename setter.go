package bvmgo_reflect

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

// SetValue function assigns the value to the target pointer.
//
// SetValue function returns an error if:
//   - targetPointer is not a pointer,
//   - value type is incompatible with targetPointer type.
func SetValue[T any](targetPointer any, value T) error {
	ptrTarget := reflect.ValueOf(targetPointer)
	// Check is not null
	if !ptrTarget.IsValid() || (ptrTarget.Kind() == reflect.Ptr && ptrTarget.IsNil()) {
		return fmt.Errorf("a not nil pointer is required to set value")
	}
	// Check is a pointer
	if ptrTarget.Kind() != reflect.Ptr {
		return fmt.Errorf("unsupported type [%s], a pointer is required to set value", typeName(ptrTarget.Type()))
	}

	return setValueToReflectValue(ptrTarget.Elem(), value)
}

// SetField function assigns the value to the field of structure pointer.
//
// SetField function returns an error if:
//   - targetPointer is not a pointer,
//   - fieldName is not a valid field name,
//   - fieldName is not found on targetPointer structure,
//   - value type is incompatible with structure field type.
func SetField[T any](targetStructurePointer any, fieldName string, value T) error {
	cleanFieldName := strings.TrimSpace(fieldName)
	ptrTarget := reflect.ValueOf(targetStructurePointer)
	// Check is not null
	if !ptrTarget.IsValid() || (ptrTarget.Kind() == reflect.Ptr && ptrTarget.IsNil()) {
		return fmt.Errorf("a not nil pointer is required to set value to [%s] field",
			cleanFieldName)
	}
	// Check is a pointer
	if ptrTarget.Kind() != reflect.Ptr {
		return fmt.Errorf("unsupported type [%s], a pointer to a structure is required to set value to [%s] field",
			typeName(ptrTarget.Type()), cleanFieldName)
	}
	targetElem := ptrTarget.Elem()
	// Check is a structure
	if targetElem.Kind() != reflect.Struct {
		return fmt.Errorf("unsupported type [%s], a pointer to a structure is required to set value to [%s] field",
			typeName(ptrTarget.Type()), cleanFieldName)
	}
	// Check field name
	if len(strings.TrimSpace(cleanFieldName)) == 0 {
		return fmt.Errorf("field name is empty")
	}
	fieldValue := targetElem.FieldByName(cleanFieldName)
	// Check field exists
	if !fieldValue.IsValid() {
		return fmt.Errorf("[%s.%s] field is not found",
			typeName(targetElem.Type()), cleanFieldName)
	}
	if !fieldValue.CanSet() {
		firstRune, _ := utf8.DecodeRuneInString(cleanFieldName)
		if unicode.IsLower(firstRune) {
			return fmt.Errorf("[%s.%s] field is private",
				typeName(targetElem.Type()), cleanFieldName)
		} else {
			return fmt.Errorf("[%s.%s] field is read only",
				typeName(targetElem.Type()), cleanFieldName)
		}
	}
	if err := setValueToReflectValue(fieldValue, value); err != nil {
		return fmt.Errorf("[%s.%s] field cannot be set with current value: %w",
			typeName(targetElem.Type()), cleanFieldName, err)
	}
	return nil
}

// setValueToReflectValue function assigns the value to the target.
//
// setValueToReflectValue function returns an error if:
//   - target is nil or invalid,
//   - value type is incompatible with structure field type.
func setValueToReflectValue[T any](target reflect.Value, value T) error {
	if reflect.TypeOf(value) == nil {
		// nil : set default "zero" value
		target.Set(reflect.Zero(target.Type()))
	} else {
		// Adapt element and value types to match (check pointer)
		eltType, valueType, match := findMatchType(target, reflect.ValueOf(value))
		if !match {
			return fmt.Errorf("value type [%s] is not assignable to variable type [%s]",
				typeName(target.Type()), TypeName(value))
		}
		eltType.Set(valueType)
	}
	return nil
}

// findMatchType find types for target and value.
//
// Types much verified valueType.Type().AssignableTo(target.Type()).
func findMatchType(target reflect.Value, valueType reflect.Value) (eltType reflect.Value, valType reflect.Value, match bool) {
	eltType = target
	valType = valueType
	match = false
	// Value type is assignable to element type : return types
	if valType.Type().AssignableTo(eltType.Type()) {
		match = true
	} else if eltType.Kind() != reflect.Ptr && valType.Kind() == reflect.Ptr &&
		valType.Elem().Type().AssignableTo(eltType.Type()) {
		// Remove value pointer
		valType = valType.Elem()
		match = true
	} else if eltType.Kind() == reflect.Ptr && valType.Kind() != reflect.Ptr &&
		valType.Type().AssignableTo(eltType.Type().Elem()) {
		// Convert value to value pointer
		if eltType.IsNil() {
			// Create value before assignation
			eltType.Set(reflect.New(eltType.Type().Elem()))
		}
		eltType = eltType.Elem()
		match = true
	}
	return
}
