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

	shrinkedStructs := make([]*downStruct, 0)
	value, next := shrink()
	for next {
		shrinkedStruct, ok := value.(*downStruct)
		if !ok {
			t.Errorf("Invalid shrinked value: %#v", value)
		}
		shrinkedStructs = append(shrinkedStructs, shrinkedStruct)
		value, next = shrink()
	}

	expected := []*downStruct{
		&downStruct{a: 0, b: "abcd", c: false},
		&downStruct{a: 5, b: "abcd", c: false},
		&downStruct{a: -5, b: "abcd", c: false},
		&downStruct{a: 8, b: "abcd", c: false},
		&downStruct{a: -8, b: "abcd", c: false},
		&downStruct{a: 9, b: "abcd", c: false},
		&downStruct{a: -9, b: "abcd", c: false},
		&downStruct{a: 10, b: "cd", c: false},
		&downStruct{a: 10, b: "ab", c: false},
		&downStruct{a: 10, b: "bcd", c: false},
		&downStruct{a: 10, b: "acd", c: false},
		&downStruct{a: 10, b: "abd", c: false},
		&downStruct{a: 10, b: "abc", c: false},
	}
	if !reflect.DeepEqual(shrinkedStructs, expected) {
		t.Errorf("%v does not equal %v", shrinkedStructs, expected)
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
	for i := 0; i < 100; i++ {
		result := gen(parameters)
		_, ok := result.Retrieve()
		if !ok {
			hasNoValue = true
			break
		}
	}
	if !hasNoValue {
		t.Error("Sieve is not applied")
	}

	sieve := gen(parameters).Sieve

	if !sieve(&downStruct{a: 2, b: "something", c: false}) {
		t.Error("Sieve did not pass even")
	}

	if sieve(&downStruct{a: 3, b: "something", c: false}) {
		t.Error("Sieve did pass odd")
	}
}
