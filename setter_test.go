package bvmgo_reflect

import (
	"maps"
	"reflect"
	"slices"
	"strings"
	"testing"
)

type testSetStruct struct {
	FieldString  string
	FieldInt32   int32
	FieldInt64   int64
	FieldFloat32 float32
	FieldFloat64 float64
	FieldBool    bool
	FieldSlice   []int32
	FieldArray   [4]string
	FieldMap     map[string]float32
	FieldFunc    func(int, int) int
	FieldStruct  testSetSubStruct
	FieldPointer *testSetSubStruct
	privateField string
}

type testSetSubStruct struct {
	Field1 int
}

func TestSetValue_bool(t *testing.T) {
	variable := false
	if err := SetValue(&variable, true); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if !variable {
		t.Errorf("SetValue() variable = %v, want = %v", variable, true)
	}
}

func TestSetValue_int32(t *testing.T) {
	variable := int32(1234)
	value := int32(12345)
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if variable != value {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_int64(t *testing.T) {
	variable := int64(123456)
	value := int64(1234567)
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if variable != value {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_float32(t *testing.T) {
	variable := float32(12.34)
	value := float32(123.45)
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if variable != value {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_float64(t *testing.T) {
	variable := float64(1234.56)
	value := float64(12345.67)
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if variable != value {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_string(t *testing.T) {
	variable := "test-before"
	value := "test-after"
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if variable != value {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_array(t *testing.T) {
	variable := [3]string{"123", "456", "789"}
	value := [3]string{"abc", "def", "ghi"}
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if variable != value {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_slice(t *testing.T) {
	variable := make([]string, 0)
	value := []string{"abc", "def", "ghi"}
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if !slices.Equal(variable, value) {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_sliceWithNilTarget(t *testing.T) {
	var variable []string = nil
	value := []string{"abc", "def", "ghi"}
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if !slices.Equal(variable, value) {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_map(t *testing.T) {
	variable := make(map[string]string)
	value := make(map[string]string)
	value["key1"] = "abc"
	value["key2"] = "def"
	value["key3"] = "ghi"
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if !maps.Equal(variable, value) {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_mapWithNilTarget(t *testing.T) {
	var variable map[string]string = nil
	value := make(map[string]string)
	value["key1"] = "abc"
	value["key2"] = "def"
	value["key3"] = "ghi"
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if !maps.Equal(variable, value) {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_struct(t *testing.T) {
	variable := testSetStruct{}
	value := testSetStruct{FieldString: "test", FieldInt32: 123, FieldInt64: 456}
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if !reflect.DeepEqual(variable, value) {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_pointer(t *testing.T) {
	var variable *testSetStruct = nil
	value := testSetStruct{FieldString: "test", FieldInt32: 123, FieldInt64: 456}
	if err := SetValue(&variable, value); err != nil {
		t.Errorf("SetValue() error = %v, want no Error", err)
		return
	}
	if !reflect.DeepEqual(variable, &value) {
		t.Errorf("SetValue() variable = %v, want = %v", variable, value)
	}
}

func TestSetValue_nilPointer(t *testing.T) {
	value := testSetStruct{FieldString: "test", FieldInt32: 123, FieldInt64: 456}
	err := SetValue(nil, value)
	if err == nil {
		t.Errorf("SetValue() returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "a not nil pointer is required") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "a not nil pointer is required")
	}
}

func TestSetValue_notPointer(t *testing.T) {
	variable := ""
	value := testSetStruct{}
	err := SetValue(variable, value)
	if err == nil {
		t.Errorf("SetValue() returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "a pointer is required to set value") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "a pointer is required to set value")
	}
}

func TestSetValue_badType(t *testing.T) {
	variable := ""
	value := testSetStruct{}
	err := SetValue(&variable, value)
	if err == nil {
		t.Errorf("SetValue() returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "is not assignable to variable type") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "is not assignable to variable type")
	}
}

func TestSetField_string(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldString", "new value")
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldString != "new value" {
		t.Errorf("testStruct.FieldString = [%v], want [%v]", testStruct.FieldString, "new value")
	}
}

func TestSetField_int32(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldInt32", int32(54321))
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldInt32 != 54321 {
		t.Errorf("testStruct.FieldInt32 = [%v], want [%v]", testStruct.FieldInt32, 54321)
	}
}

func TestSetField_int64(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldInt64", int64(987654321))
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldInt64 != 987654321 {
		t.Errorf("testStruct.FieldInt64 = [%v], want [%v]", testStruct.FieldInt64, 987654321)
	}
}

func TestSetField_float32(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldFloat32", float32(543.21))
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldFloat32 != 543.21 {
		t.Errorf("testStruct.FieldFloat32 = [%v], want [%v]", testStruct.FieldFloat32, 543.21)
	}
}

func TestSetField_float64(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldFloat64", float64(98765.4321))
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldFloat64 != 98765.4321 {
		t.Errorf("testStruct.FieldFloat64 = [%v], want [%v]", testStruct.FieldFloat64, 98765.4321)
	}
}

func TestSetField_bool(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldBool", true)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if !testStruct.FieldBool {
		t.Errorf("testStruct.FieldBool = [%v], want [%v]", testStruct.FieldBool, true)
	}
}

func TestSetField_slice(t *testing.T) {
	testStruct := testSetStruct{}
	testSlice := []int32{1, 2, 3, 4, 5}
	err := SetField(&testStruct, "FieldSlice", testSlice)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if !slices.Equal(testStruct.FieldSlice, testSlice) {
		t.Errorf("testStruct.FieldSlice = [%v], want [%v]", testStruct.FieldSlice, testSlice)
	}
}

func TestSetField_array(t *testing.T) {
	testStruct := testSetStruct{}
	testArray := [4]string{"11", "22", "33", "44"}
	err := SetField(&testStruct, "FieldArray", testArray)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldArray != testArray {
		t.Errorf("testStruct.FieldArray = [%v], want [%v]", testStruct.FieldArray, testArray)
	}
}

func TestSetField_map(t *testing.T) {
	testStruct := testSetStruct{}
	testMap := make(map[string]float32)

	err := SetField(&testStruct, "FieldMap", testMap)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if !maps.Equal(testStruct.FieldMap, testMap) {
		t.Errorf("testStruct.Field5 = [%v], want [%v]", testStruct.FieldMap, testMap)
	}
}

func TestSetField_func(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldFunc", func(a, b int) int { return a + b })
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldFunc(5, 10) != 15 {
		t.Errorf("testStruct.FieldFunc(5, 10) = [%v], want [%v]", testStruct.FieldFunc(5, 10), 15)
	}
}

func TestSetField_struct(t *testing.T) {
	testStruct := testSetStruct{}
	testSubStruct := testSetSubStruct{Field1: 1234}
	err := SetField(&testStruct, "FieldStruct", testSubStruct)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldStruct != testSubStruct {
		t.Errorf("testStruct.FieldBool = [%v], want [%v]", testStruct.FieldStruct, testSubStruct)
	}
}

func TestSetField_structFromPointer(t *testing.T) {
	testStruct := testSetStruct{}
	testSubStruct := testSetSubStruct{Field1: 1234}
	err := SetField(&testStruct, "FieldStruct", &testSubStruct)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldStruct != testSubStruct {
		t.Errorf("testStruct.FieldBool = [%v], want [%v]", testStruct.FieldStruct, testSubStruct)
	}
}

func TestSetField_pointer(t *testing.T) {
	testStruct := testSetStruct{}
	testSubStruct := testSetSubStruct{Field1: 1234}
	err := SetField(&testStruct, "FieldPointer", &testSubStruct)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldPointer != &testSubStruct {
		t.Errorf("testStruct.FieldBool = [%v], want [%v]", testStruct.FieldPointer, &testSubStruct)
	}
}

func TestSetField_nilPointer(t *testing.T) {
	testStruct := testSetStruct{}
	testStruct.FieldPointer = &testSetSubStruct{}
	var testSubStruct *testSetSubStruct = nil
	err := SetField(&testStruct, "FieldPointer", testSubStruct)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if testStruct.FieldPointer != nil {
		t.Errorf("testStruct.FieldBool = [%v], want [%v]", testStruct.FieldPointer, nil)
	}
}

func TestSetField_pointerFromStruct(t *testing.T) {
	testStruct := testSetStruct{}
	testSubStruct := testSetSubStruct{Field1: 1234}
	err := SetField(&testStruct, "FieldPointer", testSubStruct)
	if err != nil {
		t.Errorf("SetField(...) returns \"%v\" error, want nil (no error)", err)
		return
	}
	if *testStruct.FieldPointer != testSubStruct {
		t.Errorf("testStruct.FieldBool = [%v], want [%v]", testStruct.FieldPointer, &testSubStruct)
	}
}

func TestSetField_notPointer(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(testStruct, "FieldString", 1456.9832)
	if err == nil {
		t.Errorf("SetField(...) returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "a pointer to a structure is required to set value") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "a pointer to a structure is required to set value")
	}
}

func TestSetField_badType(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "FieldString", 1456.9832)
	if err == nil {
		t.Errorf("SetField(...) returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "field cannot be set with current value") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "field cannot be set with current value")
	}
}

func TestSetField_privateField(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "privateField", "new value")
	if err == nil {
		t.Errorf("SetField(...) returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "field is private") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "field is private")
	}
}

func TestSetField_fieldNotFound(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "fieldNotFound", "new value")
	if err == nil {
		t.Errorf("SetField(...) returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "field is not found") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "field is not found")
	}
}

func TestSetField_emptyFieldName(t *testing.T) {
	testStruct := testSetStruct{}
	err := SetField(&testStruct, "  ", "new value")
	if err == nil {
		t.Errorf("SetField(...) returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "field name is empty") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "field name is empty")
	}
}

func TestSetField_nilTarget(t *testing.T) {
	var testStruct *testSetStruct = nil
	err := SetField(testStruct, "FieldString", "new value")
	if err == nil {
		t.Errorf("SetField(...) returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "a not nil pointer is required") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "a not nil pointer is required")
	}
}

func TestSetField_nil(t *testing.T) {
	testStruct := testSetStruct{}
	_ = SetField[interface{}](&testStruct, "FieldString", nil)
	if testStruct.FieldString != "" {
		t.Errorf("testStruct.field1 = [%v], want [%v]", testStruct.FieldString, "")
	}
	_ = SetField[interface{}](&testStruct, "FieldInt32", nil)
	if testStruct.FieldInt32 != 0 {
		t.Errorf("testStruct.FieldInt32 = [%v], want [%v]", testStruct.FieldInt32, 0)
	}
	_ = SetField[interface{}](&testStruct, "FieldInt64", nil)
	if testStruct.FieldInt64 != 0 {
		t.Errorf("testStruct.FieldInt64 = [%v], want [%v]", testStruct.FieldInt64, 0)
	}
	_ = SetField[interface{}](&testStruct, "FieldFloat32", nil)
	if testStruct.FieldFloat32 != 0 {
		t.Errorf("testStruct.FieldFloat32 = [%v], want [%v]", testStruct.FieldFloat32, 0)
	}
	_ = SetField[interface{}](&testStruct, "FieldFloat64", nil)
	if testStruct.FieldFloat64 != 0 {
		t.Errorf("testStruct.FieldFloat64 = [%v], want [%v]", testStruct.FieldFloat64, 0)
	}
	_ = SetField[interface{}](&testStruct, "FieldBool", nil)
	if testStruct.FieldBool {
		t.Errorf("testStruct.FieldBool = [%v], want [%v]", testStruct.FieldBool, false)
	}
	_ = SetField[interface{}](&testStruct, "FieldSlice", nil)
	if testStruct.FieldSlice != nil {
		t.Errorf("testStruct.FieldSlice = [%v], want [%v]", testStruct.FieldSlice, nil)
	}
	_ = SetField[interface{}](&testStruct, "FieldArray", nil)
	var emptyArray [4]string
	if testStruct.FieldArray != emptyArray {
		t.Errorf("testStruct.FieldArray = [%v], want [%v]", testStruct.FieldArray, emptyArray)
	}
	_ = SetField[interface{}](&testStruct, "FieldMap", nil)
	if testStruct.FieldMap != nil {
		t.Errorf("testStruct.FieldMap = [%v], want [%v]", testStruct.FieldMap, nil)
	}
	_ = SetField[interface{}](&testStruct, "FieldFunc", nil)
	if testStruct.FieldFunc != nil {
		t.Errorf("testStruct.FieldFunc = a function, want no function")
	}
	_ = SetField[interface{}](&testStruct, "FieldStruct", nil)
	emptySubStruct := testSetSubStruct{}
	if testStruct.FieldStruct != emptySubStruct {
		t.Errorf("testStruct.FieldStruct = [%v], want [%v]", testStruct.FieldStruct, emptySubStruct)
	}
	_ = SetField[interface{}](&testStruct, "FieldPointer", nil)
	if testStruct.FieldPointer != nil {
		t.Errorf("testStruct.FieldPointer = [%v], want [%v]", testStruct.FieldPointer, nil)
	}
}
