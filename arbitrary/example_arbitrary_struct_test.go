package arbitrary_test

import (
	"fmt"
	"time"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
)

type MyStringType string
type MyInt8Type int8
type MyInt16Type int16
type MyInt32Type int32
type MyInt64Type int64
type MyUInt8Type uint8
type MyUInt16Type uint16
type MyUInt32Type uint32
type MyUInt64Type uint64

type Foo struct {
	Name     MyStringType
	Id1      MyInt8Type
	Id2      MyInt16Type
	Id3      MyInt32Type
	Id4      MyInt64Type
	Id5      MyUInt8Type
	Id6      MyUInt16Type
	Id7      MyUInt32Type
	Id8      MyUInt64Type
	ATime    time.Time
	ATimePtr *time.Time
}

func (f Foo) ToString() string {
	return fmt.Sprintf("For(%s, %d, %d, %d, %d, %d, %d, %d, %d, %v, %v)", f.Name, f.Id1, f.Id2, f.Id3, f.Id4, f.Id5, f.Id6, f.Id7, f.Id8, f.ATime, f.ATimePtr)
}

func Example_arbitrary_structs() {
	time.Local = time.UTC
	parameters := gopter.DefaultTestParametersWithSeed(1234) // Example should generate reproducible results, otherwise DefaultTestParameters() will suffice

	arbitraries := arbitrary.DefaultArbitraries()

	properties := gopter.NewProperties(parameters)

	properties.Property("MyInt64", arbitraries.ForAll(
		func(id MyInt64Type) bool {
			return id > -1000
		}))
	properties.Property("MyUInt32Type", arbitraries.ForAll(
		func(id MyUInt32Type) bool {
			return id < 2000
		}))
	properties.Property("Foo", arbitraries.ForAll(
		func(foo *Foo) bool {
			return foo.ATime.After(time.Unix(0, 0))
		}))
	properties.Property("Foo2", arbitraries.ForAll(
		func(foo Foo) bool {
			return foo.ATimePtr == nil || foo.ATimePtr.Before(time.Unix(20000, 0))
		}))

	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// ! MyInt64: Falsified after 6 passed tests.
	// ARG_0: -1000
	// ARG_0_ORIGINAL (54 shrinks): -1601066829744837253
	// ! MyUInt32Type: Falsified after 0 passed tests.
	// ARG_0: 2000
	// ARG_0_ORIGINAL (23 shrinks): 2161922319
	// + Foo: OK, passed 100 tests.
	// ! Foo2: Falsified after 1 passed tests.
	// ARG_0: {Name: Id1:0 Id2:0 Id3:0 Id4:0 Id5:0 Id6:0 Id7:0 Id8:0
	//    ATime:1970-01-01 00:00:00 +0000 UTC ATimePtr:1970-01-01 05:33:20 +0000
	//    UTC}
	// ARG_0_ORIGINAL (40 shrinks): {Name: Id1:-67 Id2:27301 Id3:-1350752892
	//    Id4:7128486677722156226 Id5:208 Id6:28663 Id7:4178604448
	//    Id8:16360504079646654692 ATime:2239-08-20 23:46:28.063412239 +0000 UTC
	//    ATimePtr:5468-08-19 13:09:39.171622464 +0000 UTC}
}
