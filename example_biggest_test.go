package gopter_test

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func wrong_biggest(ns []int) (int, error) {
	if len(ns) == 0 {
		return 0, fmt.Errorf("slice must have at least one element")
	}
	return ns[0], nil
}

func Example_biggest() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducible results

	properties := gopter.NewProperties(parameters)
	properties.Property("Non-zero length small int slice", prop.ForAll(
		func(ns []int) bool {
			result, _ := wrong_biggest(ns)
			sort.Slice(ns, func(i, j int) bool {
				return ns[i] > ns[j]
			})
			return result == ns[0]
		},
		gen.SliceOf(gen.IntRange(0, 20),
			reflect.TypeOf(int(0))).
			SuchThat(func(v interface{}) bool {
				return len(v.([]int)) > 0
			}),
	))
	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))

	// Output:
	// ! Non-zero length small int slice: Falsified after 1 passed tests.
	// ARG_0: [0 5 0]
	// ARG_0_ORIGINAL (1 shrinks): [0 7 5]
}
