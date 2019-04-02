package gopter_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

type downStruct struct {
	a int
	b string
	c bool
}

func TestDeriveGenSingleDown(t *testing.T) {
	gen := gopter.DeriveGen(
		func(a int, b string, c bool) *downStruct {
			return &downStruct{a: a, b: b, c: c}
		},
		func(d *downStruct) (int, string, bool) {
			return d.a, d.b, d.c
		},
		gen.Int(),
		gen.AnyString(),
		gen.Bool(),
	)

	sample, ok := gen.Sample()
	if !ok {
		t.Error("Sample not ok")
	}
	_, ok = sample.(*downStruct)
	if !ok {
		t.Errorf("%#v is not a downStruct", sample)
	}

	shrinker := gen(gopter.DefaultGenParameters()).Shrinker
	shrink := shrinker(&downStruct{a: 10, b: "abcd", c: false})

	shrunkStructs := make([]*downStruct, 0)
	value, next := shrink()
	for next {
		shrunkStruct, ok := value.(*downStruct)
		if !ok {
			t.Errorf("Invalid shrunk value: %#v", value)
		}
		shrunkStructs = append(shrunkStructs, shrunkStruct)
		value, next = shrink()
	}

	expected := []*downStruct{
		{a: 0, b: "abcd", c: false},
		{a: 5, b: "abcd", c: false},
		{a: -5, b: "abcd", c: false},
		{a: 8, b: "abcd", c: false},
		{a: -8, b: "abcd", c: false},
		{a: 9, b: "abcd", c: false},
		{a: -9, b: "abcd", c: false},
		{a: 10, b: "cd", c: false},
		{a: 10, b: "ab", c: false},
		{a: 10, b: "bcd", c: false},
		{a: 10, b: "acd", c: false},
		{a: 10, b: "abd", c: false},
		{a: 10, b: "abc", c: false},
	}
	if !reflect.DeepEqual(shrunkStructs, expected) {
		t.Errorf("%v does not equal %v", shrunkStructs, expected)
	}
}

func TestDeriveGenSingleDownWithSieves(t *testing.T) {
	gen := gopter.DeriveGen(
		func(a int, b string, c bool) *downStruct {
			return &downStruct{a: a, b: b, c: c}
		},
		func(d *downStruct) (int, string, bool) {
			return d.a, d.b, d.c
		},
		gen.Int().SuchThat(func(i int) bool {
			return i%2 == 0
		}),
		gen.AnyString(),
		gen.Bool(),
	)

	parameters := gopter.DefaultGenParameters()
	parameters.Rng.Seed(1234)

	hasNoValue := false
	sawEven := false
	sawOdd := false
	for i := 0; i < 100; i++ {
		result := gen(parameters)
		val, ok := result.Retrieve()
		if ok {
			ds := val.(*downStruct)
			if ds.a%2 == 0 {
				sawEven = true
			} else {
				sawOdd = true
			}
		} else {
			hasNoValue = true
		}
	}
	if !hasNoValue {
		t.Error("Sieve is not applied")
	}

	if !sawEven {
		t.Error("Sieve did not pass even")
	}

	if sawOdd {
		t.Error("Sieve did pass odd")
	}
}

func TestDeriveGenMultiDown(t *testing.T) {
	gen := gopter.DeriveGen(
		func(a int, b string, c bool, d int32) (*downStruct, int64) {
			return &downStruct{a: a, b: b, c: c}, int64(a) + int64(d)
		},
		func(d *downStruct, diff int64) (int, string, bool, int32) {
			return d.a, d.b, d.c, int32(diff - int64(d.a))
		},
		gen.Int(),
		gen.AnyString(),
		gen.Bool(),
		gen.Int32(),
	)

	sample, ok := gen.Sample()
	if !ok {
		t.Error("Sample not ok")
	}
	values, ok := sample.([]interface{})
	if !ok || len(values) != 2 {
		t.Errorf("%#v is not a slice of interface", sample)
	}
	_, ok = values[0].(*downStruct)
	if !ok {
		t.Errorf("%#v is not a downStruct", values[0])
	}
	_, ok = values[1].(int64)
	if !ok {
		t.Errorf("%#v is not a int64", values[1])
	}

	shrinker := gen(gopter.DefaultGenParameters()).Shrinker
	shrink := shrinker([]interface{}{&downStruct{a: 10, b: "abcd", c: false}, int64(20)})

	value, next := shrink()
	shrunkValues := make([][]interface{}, 0)
	for next {
		shrunk, ok := value.([]interface{})
		if !ok || len(values) != 2 {
			t.Errorf("%#v is not a slice of interface", sample)
		}
		shrunkValues = append(shrunkValues, shrunk)
		value, next = shrink()
	}

	expected := [][]interface{}{
		{&downStruct{a: 0, b: "abcd", c: false}, int64(10)},
		{&downStruct{a: 5, b: "abcd", c: false}, int64(15)},
		{&downStruct{a: -5, b: "abcd", c: false}, int64(5)},
		{&downStruct{a: 8, b: "abcd", c: false}, int64(18)},
		{&downStruct{a: -8, b: "abcd", c: false}, int64(2)},
		{&downStruct{a: 9, b: "abcd", c: false}, int64(19)},
		{&downStruct{a: -9, b: "abcd", c: false}, int64(1)},
		{&downStruct{a: 10, b: "cd", c: false}, int64(20)},
		{&downStruct{a: 10, b: "ab", c: false}, int64(20)},
		{&downStruct{a: 10, b: "bcd", c: false}, int64(20)},
		{&downStruct{a: 10, b: "acd", c: false}, int64(20)},
		{&downStruct{a: 10, b: "abd", c: false}, int64(20)},
		{&downStruct{a: 10, b: "abc", c: false}, int64(20)},
		{&downStruct{a: 10, b: "abcd", c: false}, int64(10)},
		{&downStruct{a: 10, b: "abcd", c: false}, int64(15)},
		{&downStruct{a: 10, b: "abcd", c: false}, int64(5)},
		{&downStruct{a: 10, b: "abcd", c: false}, int64(18)},
		{&downStruct{a: 10, b: "abcd", c: false}, int64(2)},
		{&downStruct{a: 10, b: "abcd", c: false}, int64(19)},
		{&downStruct{a: 10, b: "abcd", c: false}, int64(1)},
	}

	if !reflect.DeepEqual(shrunkValues, expected) {
		t.Errorf("%v does not equal %v", shrunkValues, expected)
	}
}

func TestDeriveGenVaryingSieveAndShrinker(t *testing.T) {
	gen := gopter.DeriveGen(
		func(a interface{}) interface{} {
			return a
		},
		func(a interface{}) interface{} {
			return a
		},
		gen.OneGenOf(gen.AnyString(), gen.Int()),
	)

	parameters := gopter.DefaultGenParameters()
	parameters.Rng.Seed(1234)

	for i := 0; i < 20; i++ {
		result := gen(parameters)
		sample, ok := result.Retrieve()
		if !ok {
			t.Error("Sample not ok")
		}
		if stringval, ok := sample.(string); ok {
			// check that the Shrinker doesn't panic
			result.Shrinker(stringval)
		} else if intval, ok := sample.(int); ok {
			result.Shrinker(intval)
		} else {
			t.Errorf("%#v is not a string or int", sample)
		}
	}
}
