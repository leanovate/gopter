package convey

import (
	"bytes"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	"github.com/leanovate/gopter/prop"
)

// ShouldSucceedForAll checks that a check condition is be true for all values, if the
// condition falsiies the generated values will be shrunk.
//
// "condition" has to be a function with the same number of parameters as the provided
// generators "gens". The function may return a simple bool (true means that the
// condition has passed), a string (empty string means that condition has passed),
// a *PropResult, or one of former combined with an error.
func ShouldSucceedForAll(condition interface{}, params ...interface{}) string {
	var arbitraries *arbitrary.Arbitraries
	parameters := gopter.DefaultTestParameters()
	gens := make([]gopter.Gen, 0)
	for _, param := range params {
		switch param.(type) {
		case *arbitrary.Arbitraries:
			arbitraries = param.(*arbitrary.Arbitraries)
		case *gopter.TestParameters:
			parameters = param.(*gopter.TestParameters)
		case gopter.Gen:
			gens = append(gens, param.(gopter.Gen))
		}
	}

	var property gopter.Prop
	if arbitraries != nil {
		property = arbitraries.ForAll(condition)
	} else {
		property = prop.ForAll(condition, gens...)
	}
	result := property.Check(parameters)

	if !result.Passed() {
		buffer := bytes.NewBufferString("")
		reporter := gopter.NewFormatedReporter(true, 75, buffer)
		reporter.ReportTestResult("", result)

		return buffer.String()
	}
	return ""
}
