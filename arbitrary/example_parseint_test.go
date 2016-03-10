package arbitrary_test

import (
	"fmt"
	"strconv"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
)

func Example_parseint() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results

	arbitraries := arbitrary.DefaultArbitraries()
	properties := gopter.NewProperties(parameters)

	properties.Property("printed integers can be parsed", arbitraries.ForAll(
		func(a int64) bool {
			str := fmt.Sprintf("%d", a)
			parsed, err := strconv.ParseInt(str, 10, 64)
			return err == nil && parsed == a
		}))

	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + printed integers can be parsed: OK, passed 100 tests.
}
