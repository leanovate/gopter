package prop_test

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func Example_shrink() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results

	properties := gopter.NewProperties(parameters)

	properties.Property("fail above 100", prop.ForAll1(
		gen.Int64(),
		func(arg interface{}) (interface{}, error) {
			return arg.(int64) <= 100, nil
		},
	))

	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// ! fail above 100: Falsified after 0 passed tests.
	// ARG_0: 101
	// ARG_0_ORIGINAL (56 shrinks): 2041104533947223744
}
