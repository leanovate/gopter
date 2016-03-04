package arbitrary_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter/arbitrary"
)

func TestArbitrariesSimple(t *testing.T) {
	arbitraries := arbitrary.DefaultArbitraries()

	gen := arbitraries.GenForType(reflect.TypeOf(true))
	value, ok := gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(bool); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(0))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(int); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(uint(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(uint); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(int8(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(int8); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(uint8(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(uint8); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(int16(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(int16); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(uint16(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(uint16); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(int32(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(int32); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(uint32(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(uint32); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(int64(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(int64); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(uint64(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(uint64); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(float32(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(float32); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(float64(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(float64); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(complex128(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(complex128); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(complex64(0)))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(complex64); !ok {
		t.Errorf("Invalid value %#v", value)
	}

	gen = arbitraries.GenForType(reflect.TypeOf(""))
	value, ok = gen.Sample()
	if !ok {
		t.Errorf("Invalid value %#v", value)
	}
	if _, ok = value.(string); !ok {
		t.Errorf("Invalid value %#v", value)
	}
}
