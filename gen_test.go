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
	var mappedWith string
	mapper := func(v string) string {
		mappedWith = v
		return "other"
	}
	value, ok := gen.Map(mapper).Sample()
	if !ok || value != "other" {
		t.Errorf("Invalid gen sample: %#v", value)
	}
	if mappedWith != "sample" {
		t.Errorf("Invalid mapped with: %#v", mappedWith)
	}

	gen = gen.SuchThat(func(interface{}) bool {
		return false
	})
	value, ok = gen.Map(mapper).Sample()
	if ok {
		t.Errorf("Invalid gen sample: %#v", value)
	}
}

func TestGenFlatMap(t *testing.T) {
	gen := constGen("sample")
	var mappedWith interface{}
	mapper := func(v interface{}) gopter.Gen {
		mappedWith = v
		return constGen("other")
	}
	value, ok := gen.FlatMap(mapper, reflect.TypeOf("")).Sample()
	if !ok || value != "other" {
		t.Errorf("Invalid gen sample: %#v", value)
	}
	if mappedWith.(string) != "sample" {
		t.Errorf("Invalid mapped with: %#v", mappedWith)
	}

	gen = gen.SuchThat(func(interface{}) bool {
		return false
	})
	value, ok = gen.FlatMap(mapper, reflect.TypeOf("")).Sample()
	if ok {
		t.Errorf("Invalid gen sample: %#v", value)
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

	gens[0] = gens[0].SuchThat(func(interface{}) bool {
		return false
	})
	gen = gopter.CombineGens(gens...)
	raw, ok = gen.Sample()
	if ok {
		t.Errorf("Invalid combined gen: %#v", raw)
	}
}

func TestSuchThat(t *testing.T) {
	var sieveArg interface{}
	sieve := func(v interface{}) bool {
		sieveArg = v
		return true
	}
	gen := constGen("sample").SuchThat(sieve)
	value, ok := gen.Sample()
	if !ok || value != "sample" {
		t.Errorf("Invalid result: %#v", value)
	}
	if sieveArg != "sample" {
		t.Errorf("Invalid sieveArg: %#v", sieveArg)
	}

	sieveArg = nil
	var sieve2Arg interface{}
	sieve2 := func(v interface{}) bool {
		sieve2Arg = v
		return false
	}
	gen = gen.SuchThat(sieve2)
	_, ok = gen.Sample()
	if ok {
		t.Error("Did not expect a result")
	}
	if sieveArg != "sample" {
		t.Errorf("Invalid sieveArg: %#v", sieveArg)
	}
	if sieve2Arg != "sample" {
		t.Errorf("Invalid sieve2Arg: %#v", sieve2Arg)
	}
}

func TestWithShrinker(t *testing.T) {
	var shrinkerArg interface{}
	shrinker := func(v interface{}) gopter.Shrink {
		shrinkerArg = v
		return gopter.NoShrink
	}
	gen := constGen("sample").WithShrinker(shrinker)
	result := gen(gopter.DefaultGenParameters())
	value, ok := result.Retrieve()
	if !ok {
		t.Errorf("Invalid combined value: %#v", value)
	}
	result.Shrinker(value)
	if shrinkerArg != "sample" {
		t.Errorf("Invalid shrinkerArg: %#v", shrinkerArg)
	}
}
