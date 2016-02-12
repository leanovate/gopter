package prop_test

import (
	"testing"

	"github.com/untoldwind/gopter"
	"github.com/untoldwind/gopter/gen"
	"github.com/untoldwind/gopter/prop"
)

func TestForAllNoShrink(t *testing.T) {
	parameters := gopter.DefaultCheckParameters()
	simpleForAll := prop.ForAllNoShrink(
		prop.NewCheck1(func(value interface{}) (interface{}, error) {
			return value.(string) == "const value", nil
		}), gen.Const("const value"))

	simpleResult := simpleForAll.Check(parameters)

	if simpleResult.Status != gopter.CheckPassed || simpleResult.Succeeded != parameters.MinSuccessfulTests {
		t.Errorf("Invalid simpleResult: %#v", simpleResult)
	}

	simpleForAllFail := prop.ForAllNoShrink(
		prop.NewCheck1(func(value interface{}) (interface{}, error) {
			return value.(string) != "const value", nil
		}), gen.Const("const value"))

	simpleResultFail := simpleForAllFail.Check(parameters)

	if simpleResultFail.Status != gopter.CheckFailed || simpleResultFail.Succeeded != 0 {
		t.Errorf("Invalid simpleResultFail: %#v", simpleResultFail)
	}
}
