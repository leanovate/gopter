package prop_test

import (
	"math"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestSqrt(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("greater one of all greater one", prop.ForAll(
		func(v float64) bool {
			return math.Sqrt(v) >= 1
		},
		gen.Float64Range(1, math.MaxFloat64),
	))

	properties.Property("greater one of all greater one alternative", prop.ForAll1(
		gen.Float64Range(1, math.MaxFloat64),
		func(v interface{}) (interface{}, error) {
			return math.Sqrt(v.(float64)) >= 1, nil
		},
	))

	properties.Property("squared is equal to value", prop.ForAll(
		func(v float64) bool {
			r := math.Sqrt(v)
			return math.Abs(r*r-v) < 1e-10*v
		},
		gen.Float64Range(0, math.MaxFloat64),
	))

	properties.Property("squared is equal to value alternative", prop.ForAll1(
		gen.Float64Range(0, math.MaxFloat64),
		func(v interface{}) (interface{}, error) {
			s := v.(float64)
			r := math.Sqrt(s)
			return math.Abs(r*r-s) < 1e-10*s, nil
		},
	))

	properties.TestingRun(t)
}
