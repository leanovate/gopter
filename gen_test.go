package gopter_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
)

func constGen(value interface{}) gopter.Gen {
	return func(*gopter.GenParameters) *gopter.GenResult {
		return gopter.NewGenResult(value, gopter.NoShrinker)
	}
}

func TestGenSample(t *testing.T) {
	gen := constGen("sample")

	value, ok := gen.Sample()
	if !ok || value != "sample" {
		t.Errorf("Invalid gen sample: %#v", value)
	}
}

func TestGenMap(t *testing.T) {
	gen := constGen("sample")
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

func TestGenFlatMap(t *testing.T) {
	gen := constGen("sample")
	var mappedWith interface{}
	mapper := func(v interface{}) gopter.Gen {
		mappedWith = v
		return constGen("other")
	}
	value, ok := gen.FlatMap(mapper).Sample()
	if !ok || value != "other" {
		t.Errorf("Invalid gen sample: %#v", value)
	}
	if mappedWith.(string) != "sample" {
		t.Errorf("Invalid mapped with: %#v", mappedWith)
	}
}

func TestCombineGens(t *testing.T) {
	gens := make([]gopter.Gen, 0, 20)
	for i := 0; i < 20; i++ {
		gens = append(gens, constGen(i))
	}
	gen := gopter.CombineGens(gens...)
	raw, ok := gen.Sample()
	if !ok {
		t.Errorf("Invalid combined gen: %#v", raw)
	}
	values, ok := raw.([]interface{})
	if !ok || !reflect.DeepEqual(values, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Errorf("Invalid combined gen: %#v", raw)
	}
}

func TestSuchThat(t *testing.T) {
	var sieveArg interface{}
	sieve := func(v interface{}) bool {
		sieveArg = v
		return false
	}
	gen := constGen("sample").SuchThat(sieve)
	_, ok := gen.Sample()
	if ok {
		t.Error("Did not expect a result")
	}
	if sieveArg != "sample" {
		t.Errorf("Invalid sieveArg: %#v", sieveArg)
	}
}
