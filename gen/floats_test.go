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

func TestFloat64Range(t *testing.T) {
	fail := gen.Float64Range(200, 100)

	if value, ok := fail.Sample(); value != nil || ok {
		t.Fail()
	}

	floats := gen.Float64Range(-1234.5, 56789.123)
	for i := 0; i < 100; i++ {
		value, ok := floats.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid floats: %#v", value)
		}
		v, ok := value.(float64)
		if !ok || math.IsNaN(v) || math.IsInf(v, 0) || v < -1234.5 || v > 56789.123 {
			t.Errorf("Invalid floats: %#v", value)
		}
	}
}

func TestFloat32(t *testing.T) {
	floats := gen.Float32()
	for i := 0; i < 100; i++ {
		value, ok := floats.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid floats: %#v", value)
		}
		_, ok = value.(float32)
		if !ok {
			t.Errorf("Invalid floats: %#v", value)
		}
	}
}

func TestFloat32Range(t *testing.T) {
	fail := gen.Float32Range(200, 100)

	if value, ok := fail.Sample(); value != nil || ok {
		t.Fail()
	}

	floats := gen.Float32Range(-1234.5, 56789.123)
	for i := 0; i < 100; i++ {
		value, ok := floats.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid floats: %#v", value)
		}
		v, ok := value.(float32)
		if !ok || v < -1234.5 || v > 56789.123 {
			t.Errorf("Invalid floats: %#v", value)
		}
	}
}
