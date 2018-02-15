package arbitrary_test

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
)

type MyIDType int32

type Foo struct {
	Name string
	Id   MyIDType
}

func Example_arbitrary_structs() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234)

	arbitraries := arbitrary.DefaultArbitraries()

	properties := gopter.NewProperties(parameters)

	properties.Property("Foo", arbitraries.ForAll(
		func(foo *Foo) bool {
			return true
		}))

	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + Foo: OK, passed 100 tests.
}
