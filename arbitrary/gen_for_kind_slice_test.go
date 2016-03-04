package arbitrary_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter/arbitrary"
)

func TestArbitrariesSlices(t *testing.T) {
	arbitraries := arbitrary.DefaultArbitraries()

	gen := arbitraries.GenForType(reflect.TypeOf([]bool{}))
	value, ok := gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.([]bool); !ok {
		t.Errorf("Invalid value %#v", value)
	}

}
