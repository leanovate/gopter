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
	Value5 *string
	Value6 interface{}
	value7 string
}

func TestStruct(t *testing.T) {
	structGen := gen.Struct(reflect.TypeOf(&testStruct{}), map[string]gopter.Gen{
		"Value1":   gen.Identifier(),
		"Value2":   gen.Int64(),
		"Value3":   gen.SliceOf(gen.Int8()),
		"NotThere": gen.AnyString(),
		"Value5":   gen.PtrOf(gen.Const("v5")),
		"Value6":   gen.AnyString(),
		"value7":   gen.AnyString(),
	})
	for i := 0; i < 100; i++ {
		value, ok := structGen.Sample()

		if !ok {
			t.Errorf("Invalid value: %#v", value)
		}
		v, ok := value.(testStruct)
		if !ok || v.Value1 == "" || v.Value3 == nil || v.Value4 != "" || !(v.Value5 == nil || *v.Value5 == "v5") {
			t.Errorf("Invalid value: %#v", value)
		}
	}
}

func TestStructWithDifferentValueTypesInSameField(t *testing.T) {
	structGen := gen.Struct(reflect.TypeOf(&testStruct{}), map[string]gopter.Gen{
		"Value6": gen.OneGenOf(gen.AnyString(), gen.Int()),
	})
	for i := 0; i < 100; i++ {
		value, ok := structGen.Sample()

		if !ok {
			t.Errorf("Invalid value: %#v", value)
		}
	}
}

func TestStructDeterminism(t *testing.T) {
	structGen := gen.Struct(reflect.TypeOf(&testStruct{}), map[string]gopter.Gen{
		"Value1": gen.Identifier(),
		"Value2": gen.Int64(),
		"Value3": gen.SliceOf(gen.Int8()),
	})
	for i := 0; i < 100; i++ {
		parameters := gopter.DefaultGenParameters().CloneWithSeed(1234)
		for _, expected := range []testStruct{
			{
				Value1: "hUeNzDbtiF4xxkidfvLaiczgpwsqfyvbbuhrjjoez4jtewulIKwzMguttazo3qwi5ufIfi6izpqT4evzrmgtmk1gQo",
				Value2: -2282921689139609493,
				Value3: []int8{-93, -96, -23, -58, 65, -108, 56, 63, -64, 26, -69, 62, 61, -93, -107, 52, -95},
			},
			{
				Value1: "ubJrJEawwnoh63jv1lxd7xhtaqqrEnjawudgiixhhkw6sdmqdgxbabyoxcoE0uviwDupccvYvxcqOv0z8opjk",
				Value2: -1611599231975617329,
				Value3: []int8{15, -41, -106, 37, 3, 76, -65, -87, 113, -115, 76, 61, 41, 65, 11, -90, -4, 43, 110, -121, 65, 112, -128, 51, -86, 50, 30, 33, -73, -88, 94, 101, 63, -113, 45, 110, 46, 21, 115, 78, -58, 47, -110, 7, -14, -18, 2, -26, 63, -33, 77, 82, -52, -57, -105},
			},
			{
				Value1: "axnbggD6Hgsxyxd6ZwcZ4Bn1uM7hzd0azvsuLvj3wvfvoramjcltivmditt5qhmHYfn0egagcFpuAffzaWxvalEaniojczez",
				Value2: -345052727922296584,
				Value3: []int8{-61, 94, 67, 9, 39, 119, 23, 1, 57, -66, 57, -94, 38, -122, 16, 82, -119, 21, -74, -66, -111, 55, -96, 8, -79, 13, -41, 124, 71, -63, 56, 16, 62, 55, -13, -35, -27, 68, -82, 22, -63, -76, 96, 60, -89, -10, -65, -102, -97, 45, 124, 117, -37, 21, 58, -87, 116, 60, -111, 27, 102, -102, -81, -123, -86, -95},
			},
			{
				Value1: "lr",
				Value2: -6442088894944465291,
				Value3: []int8{-20, 38, 25, -76, -110, -98, -61, 65, -14, 52, -47, 22, -90},
			},
			{
				Value1: "un23mggozHs4txZtydz6mIBymnxjxklkjyNzf",
				Value2: -26686468742269553,
				Value3: []int8{78, 91, -22, -126, -93, 35, -14, 67, -97, 13, -25, 73, -111, 26, 14, -67, 50, -23, -15, -63, -40, -103, 126, 60, -63, -83, -126, 64, 52, -50, 86, -25, -1, 108, 7, 62, 79, 89, 45, -73, 52, -7, -85, -111, -120, -21, 116, -8, -22, 34, 85, 36, 124, 12, -111, -114, -115, 91, -94, 82, -3, -46, 94, -73, 62, -117, -7, 84, -94, 13, 71, -5, 21, 32, 106, -44, 46},
			},
		} {
			value, ok := structGen(parameters).Retrieve()

			if !ok {
				t.Errorf("Invalid value: %#v", value)
			}
			v, ok := value.(testStruct)
			if !reflect.DeepEqual(expected, v) {
				t.Errorf("Invalid value: %#v; expected: %#v", v, expected)
			}
		}
	}
}

func TestStructPropageEmpty(t *testing.T) {
	fail := gen.Struct(reflect.TypeOf(&testStruct{}), map[string]gopter.Gen{
		"Value1": gen.Identifier().SuchThat(func(str string) bool {
			return false
		}),
	})

	if _, ok := fail.Sample(); ok {
		t.Errorf("Failing field generator in Struct generated a value")
	}
}

func TestStructNoStruct(t *testing.T) {
	fail := gen.Struct(reflect.TypeOf(""), map[string]gopter.Gen{})

	if _, ok := fail.Sample(); ok {
		t.Errorf("Invalid Struct generated a value")
	}
}

func TestStructPtr(t *testing.T) {
	structGen := gen.StructPtr(reflect.TypeOf(&testStruct{}), map[string]gopter.Gen{
		"Value1":   gen.Identifier(),
		"Value2":   gen.Int64(),
		"Value3":   gen.SliceOf(gen.Int8()),
		"NotThere": gen.AnyString(),
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

func TestStructPtrPropageEmpty(t *testing.T) {
	fail := gen.StructPtr(reflect.TypeOf(&testStruct{}), map[string]gopter.Gen{
		"Value1": gen.Identifier().SuchThat(func(str string) bool {
			return false
		}),
	})

	if _, ok := fail.Sample(); ok {
		t.Errorf("Failing field generator in StructPtr generated a value")
	}
}

func TestStructPtrNoStruct(t *testing.T) {
	fail := gen.StructPtr(reflect.TypeOf(""), map[string]gopter.Gen{})

	if _, ok := fail.Sample(); ok {
		t.Errorf("Invalid StructPtr generated a value")
	}
}
