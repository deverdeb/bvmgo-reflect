package bvmgo_reflect

import (
	"reflect"
	"testing"
	"unsafe"
)

type testTypeNameStructure struct {
	attribute int
}

func (test *testTypeNameStructure) method1(_ string) int {
	return 0
}

func (test *testTypeNameStructure) method2(function func(string) int) int {
	return function("my test")
}

func TestTypeName(t *testing.T) {
	structure := testTypeNameStructure{}
	var value = 45.6
	unsafePtr := unsafe.Pointer(&value)
	tests := []struct {
		name string
		arg  any
		want string
	}{
		{name: "nil", arg: nil, want: "<nil>"},
		{name: "int type", arg: 123, want: "int"},
		{name: "bool type", arg: true, want: "bool"},
		{name: "string type", arg: "azerty", want: "string"},
		{name: "array type", arg: [3]int{1, 2, 3}, want: "[3]int"},
		{name: "slice type", arg: make([]int, 0, 3), want: "[]int"},
		{name: "map type", arg: make(map[string]int), want: "map[string]int"},
		{name: "chan type", arg: make(chan string), want: "chan string"},
		{name: "struct type", arg: testTypeNameStructure{}, want: "bvmgo_reflect.testTypeNameStructure"},
		{name: "pointer type", arg: &testTypeNameStructure{}, want: "*bvmgo_reflect.testTypeNameStructure"},
		{name: "func type", arg: TestTypeName, want: "func(*testing.T)"},
		{name: "method type", arg: structure.method1, want: "func(string) int"},
		{name: "method func type", arg: structure.method2, want: "func(func(string) int) int"},
		{name: "unsafe pointer type", arg: unsafePtr, want: "unsafe.Pointer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeName(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
