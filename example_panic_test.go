package gopter_test

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func Example_panic() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducible results

	properties := gopter.NewProperties(parameters)
	properties.Property("Will panic", prop.ForAll(
		func(i int) bool {
			if i%2 == 0 {
				panic("hi")
			}
			return true
		},
		gen.Int().WithLabel("number")))
	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// ! Will panic: Error on property evaluation after 6 passed tests: Check
	//    paniced: hi
	// number: 0
	// number_ORIGINAL (1 shrinks): 2015020988
}
