package gen_test

import (
	"math"
	"testing"

	"github.com/leanovate/gopter/gen"
)

func TestFloat64(t *testing.T) {
	floats := gen.Float64()
	for i := 0; i < 100; i++ {
		value, ok := floats.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid floats: %#v", value)
		}
		v, ok := value.(float64)
		if !ok || math.IsNaN(v) || math.IsInf(v, 0) {
			t.Errorf("Invalid floats: %#v", value)
		}
	}
}
