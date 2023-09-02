package bvmgo_reflect

import (
	"reflect"
)

// TypeName return the name of type.
func TypeName[T any](value T) string {
	return typeName(reflect.TypeOf(value))
}

// typeName return the name of type.
func typeName(elementType reflect.Type) string {
	if elementType == nil {
		return "<nil>"
	}
	if elementType.Kind() == reflect.Invalid {
		return "<invalid>"
	}
	return elementType.String()
}
