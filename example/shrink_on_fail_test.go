package example

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestShrinkOnFail(t *testing.T) {
	parameters := gopter.DefaultTestParameters()

	failAbove100 := prop.ForAll1(
		gen.Int64(),
		func(arg interface{}) (interface{}, error) {
			return arg.(int64) <= 100, nil
		},
	)

	result := failAbove100.Check(parameters)
	if result.Passed() || len(result.Args) != 1 {
		t.Errorf("Expected to fail with args: %#v", result)
	}
	arg := result.Args[0]
	if arg.Arg.(int64) != 101 || arg.Arg.(int64) >= arg.OrigArg.(int64) {
		t.Errorf("Arg has not been shrinked: %#v", arg)
	}
}
