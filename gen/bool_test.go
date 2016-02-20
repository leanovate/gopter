package gen_test

import (
	"testing"

	"github.com/leanovate/gopter/gen"
)

func TestBool(t *testing.T) {
	b := gen.Bool()
	for i := 0; i < 100; i++ {
		value, ok := b.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid bool: %#v", value)
		}
		_, ok = value.(bool)
		if !ok {
			t.Errorf("Invalid bool: %#v", value)
		}
	}

}
