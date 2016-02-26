package gen_test

import (
	"math"
	"testing"

	"github.com/leanovate/gopter/gen"
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
			t.Errorf("Invalid limited: %#v", value)
		}
		v, ok := value.(int64)
		if !ok || v < -123456 || v > 234567 {
			t.Errorf("Invalid limited: %#v", value)
		}
	}

	pos := gen.Int64Range(1, math.MaxInt64)
	for i := 0; i < 100; i++ {
		value, ok := pos.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid pos: %#v", value)
		}
		v, ok := value.(int64)
		if !ok || v <= 0 {
			t.Errorf("Invalid pos: %#v", value)
		}
	}

	neg := gen.Int64Range(math.MinInt64, -1)
	for i := 0; i < 100; i++ {
		value, ok := neg.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid neg: %#v", value)
		}
		v, ok := value.(int64)
		if !ok || v >= 0 {
			t.Errorf("Invalid neg: %#v", value)
		}
	}
}

func TestUInt64Range(t *testing.T) {
	fail := gen.UInt64Range(200, 100)

	if value, ok := fail.Sample(); value != nil || ok {
		t.Fail()
	}

	limited := gen.UInt64Range(0, 234567)
	for i := 0; i < 100; i++ {
		value, ok := limited.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		v, ok := value.(uint64)
		if !ok || v > 234567 {
			t.Errorf("Invalid limited: %#v", value)
		}
	}
}

func TestInt64(t *testing.T) {
	ints := gen.Int64()
	for i := 0; i < 100; i++ {
		value, ok := ints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(int64)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}

	uints := gen.UInt64()
	for i := 0; i < 100; i++ {
		value, ok := uints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(uint64)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}
}

func TestInt32(t *testing.T) {
	ints := gen.Int32()
	for i := 0; i < 100; i++ {
		value, ok := ints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(int32)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}

	uints := gen.UInt32()
	for i := 0; i < 100; i++ {
		value, ok := uints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(uint32)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}
}

func TestInt16(t *testing.T) {
	ints := gen.Int16()
	for i := 0; i < 100; i++ {
		value, ok := ints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(int16)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}

	uints := gen.UInt16()
	for i := 0; i < 100; i++ {
		value, ok := uints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(uint16)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}
}

func TestInt8(t *testing.T) {
	ints := gen.Int8()
	for i := 0; i < 100; i++ {
		value, ok := ints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(int8)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}

	uints := gen.UInt8()
	for i := 0; i < 100; i++ {
		value, ok := uints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(uint8)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(byte)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}
}

func TestInt(t *testing.T) {
	ints := gen.Int()
	for i := 0; i < 100; i++ {
		value, ok := ints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(int)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}

	uints := gen.UInt()
	for i := 0; i < 100; i++ {
		value, ok := uints.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid limited: %#v", value)
		}
		_, ok = value.(uint)
		if !ok {
			t.Errorf("Invalid limited: %#v", value)
		}
	}
}
