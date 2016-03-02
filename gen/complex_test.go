package gen_test

import (
	"math"
	"testing"

	"github.com/leanovate/gopter/gen"
)

func TestComplex128Box(t *testing.T) {
	minReal := -12345.67
	maxReal := 2345.78
	minImag := -5432.8
	maxImag := 8764.6
	complexs := gen.Complex128Box(complex(minReal, minImag), complex(maxReal, maxImag))
	for i := 0; i < 100; i++ {
		value, ok := complexs.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid complexs: %#v", value)
		}
		v, ok := value.(complex128)
		if !ok || real(v) < minReal || real(v) > maxReal || imag(v) < minImag || imag(v) > maxImag {
			t.Errorf("Invalid complexs: %#v", value)
		}
	}
}

func TestComplex128(t *testing.T) {
	complexs := gen.Complex128()
	for i := 0; i < 100; i++ {
		value, ok := complexs.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid complexs: %#v", value)
		}
		v, ok := value.(complex128)
		if !ok || math.IsNaN(real(v)) || math.IsNaN(imag(v)) || math.IsInf(real(v), 0) || math.IsInf(imag(v), 0) {
			t.Errorf("Invalid complexs: %#v", value)
		}
	}
}

func TestComplex64Box(t *testing.T) {
	minReal := float32(-12345.67)
	maxReal := float32(2345.78)
	minImag := float32(-5432.8)
	maxImag := float32(8764.6)
	complexs := gen.Complex64Box(complex(minReal, minImag), complex(maxReal, maxImag))
	for i := 0; i < 100; i++ {
		value, ok := complexs.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid complexs: %#v", value)
		}
		v, ok := value.(complex64)
		if !ok || real(v) < minReal || real(v) > maxReal || imag(v) < minImag || imag(v) > maxImag {
			t.Errorf("Invalid complexs: %#v", value)
		}
	}
}

func TestComplex64(t *testing.T) {
	complexs := gen.Complex64()
	for i := 0; i < 100; i++ {
		value, ok := complexs.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid complexs: %#v", value)
		}
		_, ok = value.(complex64)
		if !ok {
			t.Errorf("Invalid complexs: %#v", value)
		}
	}
}
