package gen_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

type testStruct struct {
	Value1 string
	Value2 int64
	Value3 []int8
	Value4 string
}

func TestStructPtr(t *testing.T) {
	structGen := gen.StructPtr(reflect.TypeOf(&testStruct{}), map[string]gopter.Gen{
		"Value1": gen.Identifier(),
		"Value2": gen.Int64(),
		"Value3": gen.SliceOf(gen.Int8()),
	})
	for i := 0; i < 100; i++ {
		value, ok := structGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid value: %#v", value)
		}
		v, ok := value.(*testStruct)
		if !ok || v.Value1 == "" || v.Value3 == nil || v.Value4 != "" {
			t.Errorf("Invalid value: %#v", value)
		}
	}
}
