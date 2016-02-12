package gen_test

import (
	"math"
	"testing"

	"github.com/untoldwind/gopter/gen"
)

func TestInt64Range(t *testing.T) {
	fail := gen.Int64Range(200, 100)

	if value, ok := fail.Sample(); value != nil || ok {
		t.Fail()
	}

	limited := gen.Int64Range(-123456, 234567)
	for i := 0; i < 100; i++ {
		value, ok := limited.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: $#v", value)
		}
		v, ok := value.(int64)
		if !ok || v < -123456 || v > 234567 {
			t.Errorf("Invalid limited: $#v", value)
		}
	}

	pos := gen.Int64Range(1, math.MaxInt64)
	for i := 0; i < 100; i++ {
		value, ok := pos.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid pos: $#v", value)
		}
		v, ok := value.(int64)
		if !ok || v <= 0 {
			t.Errorf("Invalid pos: $#v", value)
		}
	}

	neg := gen.Int64Range(math.MinInt64, -1)
	for i := 0; i < 100; i++ {
		value, ok := neg.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid neg: $#v", value)
		}
		v, ok := value.(int64)
		if !ok || v >= 0 {
			t.Errorf("Invalid neg: $#v", value)
		}
	}
}
