package gen

import (
	"math"
	"reflect"

	"github.com/leanovate/gopter"
)

// Int64Range generates int64 numbers within a given range
func Int64Range(min, max int64) gopter.Gen {
	if max < min {
		return Fail(reflect.TypeOf(int64(0)))
	}
	rangeSize := uint64(max - min + 1)

	if max == math.MaxInt64 && min == math.MinInt64 { // Check overflow (i.e. max = MaxInt64, min = MinInt64)
		return func(genParams *gopter.GenParameters) *gopter.GenResult {
			return gopter.NewGenResult(genParams.NextInt64(), Int64Shrinker)
		}
	}
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		var nextResult uint64 = uint64(min) + (genParams.NextUint64() % rangeSize)
		genResult := gopter.NewGenResult(int64(nextResult), Int64Shrinker)
		genResult.Sieve = func(v interface{}) bool {
			return v.(int64) >= min && v.(int64) <= max
		}
		return genResult
	}
}

// UInt64Range generates uint64 numbers within a given range
func UInt64Range(min, max uint64) gopter.Gen {
	if max < min {
		return Fail(reflect.TypeOf(uint64(0)))
	}
	d := max - min + 1
	if d == 0 { // Check overflow (i.e. max = MaxInt64, min = MinInt64)
		return func(genParams *gopter.GenParameters) *gopter.GenResult {
			return gopter.NewGenResult(genParams.NextUint64(), UInt64Shrinker)
		}
	}
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		genResult := gopter.NewGenResult(min+genParams.NextUint64()%d, UInt64Shrinker)
		genResult.Sieve = func(v interface{}) bool {
			return v.(uint64) >= min && v.(uint64) <= max
		}
		return genResult
	}
}

// Int64 generates an arbitrary int64 number
func Int64() gopter.Gen {
	return Int64Range(math.MinInt64, math.MaxInt64)
}

// UInt64 generates an arbitrary Uint64 number
func UInt64() gopter.Gen {
	return UInt64Range(0, math.MaxUint64)
}

// Int32Range generates int32 numbers within a given range
func Int32Range(min, max int32) gopter.Gen {
	return Int64Range(int64(min), int64(max)).
		Map(int64To32).
		WithShrinker(Int32Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(int32) >= min && v.(int32) <= max
	})
}

// UInt32Range generates uint32 numbers within a given range
func UInt32Range(min, max uint32) gopter.Gen {
	return UInt64Range(uint64(min), uint64(max)).
		Map(uint64To32).
		WithShrinker(UInt32Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(uint32) >= min && v.(uint32) <= max
	})
}

// Int32 generate arbitrary int32 numbers
func Int32() gopter.Gen {
	return Int32Range(math.MinInt32, math.MaxInt32)
}

// UInt32 generate arbitrary int32 numbers
func UInt32() gopter.Gen {
	return UInt32Range(0, math.MaxUint32)
}

// Int16Range generates int16 numbers within a given range
func Int16Range(min, max int16) gopter.Gen {
	return Int64Range(int64(min), int64(max)).
		Map(int64To16).
		WithShrinker(Int16Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(int16) >= min && v.(int16) <= max
	})
}

// UInt16Range generates uint16 numbers within a given range
func UInt16Range(min, max uint16) gopter.Gen {
	return UInt64Range(uint64(min), uint64(max)).
		Map(uint64To16).
		WithShrinker(UInt16Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(uint16) >= min && v.(uint16) <= max
	})
}

// Int16 generate arbitrary int16 numbers
func Int16() gopter.Gen {
	return Int16Range(math.MinInt16, math.MaxInt16)
}

// UInt16 generate arbitrary uint16 numbers
func UInt16() gopter.Gen {
	return UInt16Range(0, math.MaxUint16)
}

// Int8Range generates int8 numbers within a given range
func Int8Range(min, max int8) gopter.Gen {
	return Int64Range(int64(min), int64(max)).
		Map(int64To8).
		WithShrinker(Int8Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(int8) >= min && v.(int8) <= max
	})
}

// UInt8Range generates uint8 numbers within a given range
func UInt8Range(min, max uint8) gopter.Gen {
	return UInt64Range(uint64(min), uint64(max)).
		Map(uint64To8).
		WithShrinker(UInt8Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(uint8) >= min && v.(uint8) <= max
	})
}

// Int8 generate arbitrary int16 numbers
func Int8() gopter.Gen {
	return Int8Range(math.MinInt8, math.MaxInt8)
}

// UInt8 generate arbitrary uint16 numbers
func UInt8() gopter.Gen {
	return UInt8Range(0, math.MaxUint8)
}

// IntRange generates int numbers within a given range
func IntRange(min, max int) gopter.Gen {
	return Int64Range(int64(min), int64(max)).
		Map(int64ToInt).
		WithShrinker(IntShrinker)
}

// Int generate arbitrary int numbers
func Int() gopter.Gen {
	return Int64Range(math.MinInt32, math.MaxInt32).
		Map(int64ToInt).
		WithShrinker(IntShrinker)
}

// UIntRange generates uint numbers within a given range
func UIntRange(min, max int) gopter.Gen {
	return UInt64Range(uint64(min), uint64(max)).
		Map(uint64ToUint).
		WithShrinker(UIntShrinker)
}

// UInt generate arbitrary uint numbers
func UInt() gopter.Gen {
	return UInt64Range(0, math.MaxUint32).
		Map(uint64ToUint).
		WithShrinker(UIntShrinker)
}

func int64To32(value interface{}) interface{} {
	return int32(value.(int64))
}

func uint64To32(value interface{}) interface{} {
	return uint32(value.(uint64))
}

func int64To16(value interface{}) interface{} {
	return int16(value.(int64))
}

func uint64To16(value interface{}) interface{} {
	return uint16(value.(uint64))
}

func int64To8(value interface{}) interface{} {
	return int8(value.(int64))
}

func uint64To8(value interface{}) interface{} {
	return uint8(value.(uint64))
}

func int64ToInt(value interface{}) interface{} {
	return int(value.(int64))
}

func uint64ToUint(value interface{}) interface{} {
	return uint(value.(uint64))
}
