package gopter_test

import (
	"math"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func Example_sqrt() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results

	properties := gopter.NewProperties(parameters)

	properties.Property("greater one of all greater one", prop.ForAll1(
		gen.Float64Range(1, math.MaxFloat64),
		func(v interface{}) (interface{}, error) {
			return math.Sqrt(v.(float64)) >= 1, nil
		},
	))

	properties.Property("squared is equal to value", prop.ForAll1(
		gen.Float64Range(0, math.MaxFloat64),
		func(v interface{}) (interface{}, error) {
			s := v.(float64)
			r := math.Sqrt(s)
			return math.Abs(r*r-s) < 1e-10*s, nil
		},
	))

	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + greater one of all greater one: OK, passed 100 tests.
	// + squared is equal to value: OK, passed 100 tests.
}
