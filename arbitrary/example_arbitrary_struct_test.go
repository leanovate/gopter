package arbitrary_test

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
)

type MyStringType string
type MyInt8Type int8
type MyInt16Type int16
type MyInt32Type int32
type MyInt64Type int32
type MyUInt8Type uint8
type MyUInt16Type uint16
type MyUInt32Type uint32
type MyUInt64Type uint32

type Foo struct {
	Name MyStringType
	Id1  MyInt8Type
	Id2  MyInt16Type
	Id3  MyInt32Type
	Id4  MyInt64Type
	Id5  MyUInt8Type
	Id6  MyUInt16Type
	Id7  MyUInt32Type
	Id8  MyUInt64Type
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
	properties.Property("Foo2", arbitraries.ForAll(
		func(foo Foo) bool {
			return true
		}))

	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + Foo: OK, passed 100 tests.
	// + Foo2: OK, passed 100 tests.
}
