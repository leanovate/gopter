package gopter_test

import (
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	"github.com/leanovate/gopter/gen"
)

type TestBook struct {
	Title   string
	Content string
}

func genTestBook() gopter.Gen {
	return gen.Struct(reflect.TypeOf(&TestBook{}), map[string]gopter.Gen{
		"Title":   gen.AlphaString(),
		"Content": gen.AlphaString(),
	})
}

type TestLibrary struct {
	Name       string
	Librarians uint8
	Books      []TestBook
}

func genTestLibrary() gopter.Gen {
	return gen.Struct(reflect.TypeOf(&TestLibrary{}), map[string]gopter.Gen{
		"Name":       gen.AlphaString(),
		"Librarians": gen.UInt8Range(1, 255),
		"Books":      gen.SliceOf(genTestBook()),
	})
}

type CityName = string
type TestCities struct {
	Libraries map[string][]TestLibrary
}

func genTestCities() gopter.Gen {
	return gen.StructPtr(reflect.TypeOf(&TestCities{}), map[string]gopter.Gen{
		"Libraries": (gen.MapOf(gen.AlphaString(), gen.SliceOf(genTestLibrary()))),
	})
}
func Example_libraries() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results
	parameters.MaxSize = 5
	arbitraries := arbitrary.DefaultArbitraries()
	arbitraries.RegisterGen(genTestCities())

	properties := gopter.NewProperties(parameters)

	properties.Property("no unsupervised libraries", arbitraries.ForAll(
		func(tc *TestCities) bool {
			for _, libraries := range tc.Libraries {
				for _, library := range libraries {
					if library.Librarians == 0 {
						return false
					}
				}
			}
			return true
		},
	))

	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + no unsupervised libraries: OK, passed 100 tests.
}
