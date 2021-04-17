package gopter_test

import (
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func Example_flatmap() {

	type IntPair struct {
		Fst int
		Snd int
	}

	// Generate a pair of integers, such that the first
	// is in the range of 10-20 and the second in the
	// in the range of 2k-50, depending on the value of
	// the first.
	genIntPair := func() gopter.Gen {
		return gen.IntRange(10, 20).FlatMap(func(v interface{}) gopter.Gen {
			k := v.(int)
			n := gen.Const(k)
			m := gen.IntRange(2*k, 50)
			var gen_map = map[string]gopter.Gen{"Fst": n, "Snd": m}
			return gen.Struct(
				reflect.TypeOf(IntPair{}),
				gen_map,
			)
		},
			reflect.TypeOf(int(0)))
	}

	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducible results
	properties := gopter.NewProperties(parameters)

	properties.Property("Generate a dependent pair of integers", prop.ForAll(
		func(p IntPair) bool {
			a := p.Fst
			b := p.Snd
			return a*2 <= b
		},
		genIntPair(),
	))

	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))

	// Output:
	// + Generate a dependent pair of integers: OK, passed 100 tests.
}
