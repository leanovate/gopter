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
		t.Error("%#v is not a downStruct", sample)
	}
}
