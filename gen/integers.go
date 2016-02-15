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
	d := uint64(max - min + 1)

	if d == 0 { // Check overflow (i.e. max = MaxInt64, min = MinInt64)
		return func(genParams *gopter.GenParameters) *gopter.GenResult {
			return gopter.NewGenResult(genParams.NextInt64(), Int64Shrinker)
		}
	}
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		genResult := gopter.NewGenResult(min+int64(genParams.NextUint64()%d), Int64Shrinker)
		genResult.Sieve = func(v interface{}) bool {
			return v.(int64) >= min && v.(int64) <= max
		}
		return genResult
	}
}

// Int64 generates an arbitrary int64 number
func Int64() gopter.Gen {
	return Int64Range(math.MinInt64, math.MaxInt64)
}

// Int32Range generates int32 numbers within a given range
func Int32Range(min, max int32) gopter.Gen {
	return Int64Range(int64(min), int64(max)).
		Map(int64To32, Int32Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(int32) >= min && v.(int32) <= max
	})
}

// Int32 generate arbitrary int32 numbers
func Int32() gopter.Gen {
	return Int32Range(math.MinInt32, math.MaxInt32)
}

// Int16Range generates int16 numbers within a given range
func Int16Range(min, max int16) gopter.Gen {
	return Int64Range(int64(min), int64(max)).
		Map(int64To16, Int16Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(int16) >= min && v.(int16) <= max
	})
}

// Int16 generate arbitrary int16 numbers
func Int16() gopter.Gen {
	return Int16Range(math.MinInt16, math.MaxInt16)
}

// Int8Range generates int8 numbers within a given range
func Int8Range(min, max int8) gopter.Gen {
	return Int64Range(int64(min), int64(max)).
		Map(int64To8, Int8Shrinker).
		SuchThat(func(v interface{}) bool {
		return v.(int8) >= min && v.(int8) <= max
	})
}

// Int8 generate arbitrary int16 numbers
func Int8() gopter.Gen {
	return Int8Range(math.MinInt8, math.MaxInt8)
}

func int64To32(value interface{}) interface{} {
	return int32(value.(int64))
}

func int64To16(value interface{}) interface{} {
	return int16(value.(int64))
}

func int64To8(value interface{}) interface{} {
	return int8(value.(int64))
}
