package gopter_test

import (
	"testing"

	"github.com/leanovate/gopter"
)

func TestGenSample(t *testing.T) {
	gen := gopter.Gen(func(*gopter.GenParameters) *gopter.GenResult {
		return gopter.NewGenResult("sample", gopter.NoShrinker)
	})

	value, ok := gen.Sample()
	if !ok || value != "sample" {
		t.Errorf("Invalid gen sample: %#v", value)
	}
}

func TestGenMap(t *testing.T) {
	gen := gopter.Gen(func(*gopter.GenParameters) *gopter.GenResult {
		return gopter.NewGenResult("sample", gopter.NoShrinker)
	})
	var mappedWith interface{}
	mapper := func(v interface{}) interface{} {
		mappedWith = v
		return "other"
	}
	value, ok := gen.Map(mapper).Sample()
	if !ok || value != "other" {
		t.Errorf("Invalid gen sample: %#v", value)
	}
	if mappedWith.(string) != "sample" {
		t.Errorf("Invalid mapped with: %#v", mappedWith)
	}
}
