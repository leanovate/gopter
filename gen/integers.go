package gen

import (
	"math"

	"github.com/untoldwind/gopter"
)

func Int64Range(min, max int64) gopter.Gen {
	if max < min {
		return Fail
	}
	d := uint64(max - min + 1)

	if d == 0 { // Check overflow (i.e. max = MaxInt64, min = MinInt64)
		return func(genParams *gopter.GenParameters) *gopter.GenResult {
			return gopter.NewGenResult(int64(genParams.NextUint64()), gopter.NoShrinker)
		}
	}
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		return gopter.NewGenResult(min+int64(genParams.NextUint64()%d), gopter.NoShrinker)
	}
}

func Int64() gopter.Gen {
	return Int64Range(math.MinInt64, math.MaxInt64)
}

func Int32Range(min, max int32) gopter.Gen {
	return Int64Range(math.MinInt32, math.MaxInt32).Map(func(value interface{}) interface{} {
		return int32(value.(int64))
	})
}

func Int16Range(min, max int16) gopter.Gen {
	return Int64Range(math.MinInt16, math.MaxInt16).Map(func(value interface{}) interface{} {
		return int16(value.(int64))
	})
}

func Int8Range(min, max int8) gopter.Gen {
	return Int64Range(math.MinInt8, math.MaxInt8).Map(func(value interface{}) interface{} {
		return int8(value.(int64))
	})
}
