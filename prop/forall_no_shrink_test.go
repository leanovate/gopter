package prop_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestForAllNoShrink(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	simpleForAll := prop.ForAllNoShrink1(
		gen.Const("const value"),
		func(value interface{}) (interface{}, error) {
			return value.(string) == "const value", nil
		},
	)

	simpleResult := simpleForAll.Check(parameters)

	if simpleResult.Status != gopter.TestPassed || simpleResult.Succeeded != parameters.MinSuccessfulTests {
		t.Errorf("Invalid simpleResult: %#v", simpleResult)
	}

	simpleForAllFail := prop.ForAllNoShrink1(
		gen.Const("const value"),
		func(value interface{}) (interface{}, error) {
			return value.(string) != "const value", nil
		},
	)

	simpleResultFail := simpleForAllFail.Check(parameters)

	if simpleResultFail.Status != gopter.TestFailed || simpleResultFail.Succeeded != 0 {
		t.Errorf("Invalid simpleResultFail: %#v", simpleResultFail)
	}
}
