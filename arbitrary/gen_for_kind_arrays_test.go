package arbitrary_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter/arbitrary"
)

func TestArbitrariesArrays(t *testing.T) {
	arbitraries := arbitrary.DefaultArbitraries()

	gen := arbitraries.GenForType(reflect.TypeOf([20]int{}))
	value, ok := gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.([20]int); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf([10]string{}))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.([10]string); !ok {
		t.Errorf("Invalid value %#v", value)
	}
}
