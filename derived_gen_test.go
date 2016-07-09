package gopter_test

import (
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

	sieve := gen(gopter.DefaultGenParameters()).Sieve

	if !sieve(&downStruct{a: 2, b: "something", c: false}) {
		t.Error("Sieve did not pass even")
	}

	if sieve(&downStruct{a: 3, b: "something", c: false}) {
		t.Error("Sieve did pass odd")
	}
}
