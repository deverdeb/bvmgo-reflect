package bvmgo_reflect

import (
	"strings"
	"testing"
)

func TestGetFieldString_struct(t *testing.T) {
	expectedValue := "test_value"
	testStruct := testSetStruct{FieldString: expectedValue}
	findValue, err := GetFieldString(testStruct, "FieldString")
	if err != nil {
		t.Errorf("GetFieldString(...) returns \"%v\" error, want no error", err)
		return
	}
	if findValue != expectedValue {
		t.Errorf("GetFieldString(...) = [%v], want [%v]", findValue, expectedValue)
	}
}

func TestGetFieldString_pointerToStruct(t *testing.T) {
	expectedValue := "test_value"
	testStruct := testSetStruct{FieldString: expectedValue}
	findValue, err := GetFieldString(&testStruct, "FieldString")
	if err != nil {
		t.Errorf("GetFieldString(...) returns \"%v\" error, want no error", err)
		return
	}
	if findValue != expectedValue {
		t.Errorf("GetFieldString(...) = [%v], want [%v]", findValue, expectedValue)
	}
}

func TestGetFieldString_map(t *testing.T) {
	expectedValue := "test_value"
	testStruct := make(map[string]string)
	testStruct["MyEntry"] = expectedValue
	findValue, err := GetFieldString(testStruct, "MyEntry")
	if err != nil {
		t.Errorf("GetFieldString(...) returns \"%v\" error, want no error", err)
		return
	}
	if findValue != expectedValue {
		t.Errorf("GetFieldString(...) = [%v], want [%v]", findValue, expectedValue)
	}
}

func TestGetFieldString_pointerToMap(t *testing.T) {
	expectedValue := "test_value"
	testStruct := make(map[string]string)
	testStruct["MyEntry"] = expectedValue
	findValue, err := GetFieldString(&testStruct, "MyEntry")
	if err != nil {
		t.Errorf("GetFieldString(...) returns \"%v\" error, want no error", err)
		return
	}
	if findValue != expectedValue {
		t.Errorf("GetFieldString(...) = [%v], want [%v]", findValue, expectedValue)
	}
}

func TestGetFieldString_nilPointer(t *testing.T) {
	var testStruct map[string]string = nil
	_, err := GetFieldString(testStruct, "MyEntry")
	if err == nil {
		t.Errorf("GetFieldString(...) returns nil (no error), want an error")
		return
	}
	if !strings.Contains(err.Error(), "a not nil pointer is required") {
		t.Errorf("err.Error() = [%v], want contain [%v]", err.Error(), "a not nil pointer is required")
	}
}
