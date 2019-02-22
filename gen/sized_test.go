package gen_test

import (
	"testing"

	"github.com/leanovate/gopter"

	"github.com/leanovate/gopter/gen"
)

func TestSized(t *testing.T) {
	sizedInt := gen.Sized(func(size int) gopter.Gen {
		return gen.IntRange(0, size)
	})

	for i := 0; i < 100; i++ {
		value, ok := sizedInt.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid sized int: %#v", value)
		}
		v, ok := value.(int)
		if !ok {
			t.Errorf("Invalid sized int: %#v", value)
		}
		if v < 0 || v > 100 {
			t.Errorf("Sized int out of range: %d", v)
		}
	}
}
