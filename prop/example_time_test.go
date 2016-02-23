package prop_test

import (
	"time"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func Example_timeGen() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results

	properties := gopter.NewProperties(parameters)

	properties.Property("regular time format parsable", prop.ForAll1(
		gen.TimeRange(time.Now(), time.Duration(100*24*365)*time.Hour),
		func(arg interface{}) (interface{}, error) {
			actual := arg.(time.Time)
			str := actual.Format(time.RFC3339Nano)
			parsed, err := time.Parse(time.RFC3339Nano, str)
			return actual.Equal(parsed), err
		},
	))

	properties.Property("any time format parsable", prop.ForAll1(
		gen.Time(),
		func(arg interface{}) (interface{}, error) {
			actual := arg.(time.Time)
			str := actual.Format(time.RFC3339Nano)
			parsed, err := time.Parse(time.RFC3339Nano, str)
			return actual.Equal(parsed), err
		},
	))

	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + regular time format parsable: OK, passed 100 tests.
	// ! any time format parsable: Error on property evaluation after 0 passed
	//    tests: parsing time "-0001-12-31T23:59:59+01:00" as
	//    "2006-01-02T15:04:05.999999999Z07:00": cannot parse
	//    "-0001-12-31T23:59:59+01:00" as "2006"
	// ARG_0: -0001-12-31 23:59:59 +0100 CET
	// ARG_0_ORIGINAL (43 shrinks): -274321993098-08-05 08:52:07.761378105 +0100
	//    CET
}
