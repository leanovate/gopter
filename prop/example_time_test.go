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
	time.Local = time.UTC     // Just for this example to generate reproducable results

	properties := gopter.NewProperties(parameters)

	properties.Property("regular time format parsable", prop.ForAll(
		func(actual time.Time) (bool, error) {
			str := actual.Format(time.RFC3339Nano)
			parsed, err := time.Parse(time.RFC3339Nano, str)
			return actual.Equal(parsed), err
		},
		gen.TimeRange(time.Now(), time.Duration(100*24*365)*time.Hour),
	))

	properties.Property("any time format parsable", prop.ForAll(
		func(actual time.Time) (bool, error) {
			str := actual.Format(time.RFC3339Nano)
			parsed, err := time.Parse(time.RFC3339Nano, str)
			return actual.Equal(parsed), err
		},
		gen.Time(),
	))

	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + regular time format parsable: OK, passed 100 tests.
	// ! any time format parsable: Error on property evaluation after 0 passed
	//    tests: parsing time "-0001-12-31T23:59:59Z" as
	//    "2006-01-02T15:04:05.999999999Z07:00": cannot parse
	//    "-0001-12-31T23:59:59Z" as "2006"
	// ARG_0: -0001-12-31 23:59:59 +0000 UTC
	// ARG_0_ORIGINAL (45 shrinks): -274321993098-08-05 07:52:07.761378105 +0000
	//    UTC
}
