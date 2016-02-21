package gen

import (
	"math"
	"reflect"

	"github.com/leanovate/gopter"
)

// Float64Range generates float64 numbers within a given range
func Float64Range(min, max float64) gopter.Gen {
	d := max - min
	if d < 0 || d > math.MaxFloat64 {
		return Fail(reflect.TypeOf(int64(0)))
	}

	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		genResult := gopter.NewGenResult(min+genParams.Rng.Float64()*d, Float64Shrinker)
		genResult.Sieve = func(v interface{}) bool {
			return v.(float64) >= min && v.(float64) <= max
		}
		return genResult
	}
}

// Float64 generates arbitrary float64 numbers that do not contain NaN or Inf
func Float64() gopter.Gen {
	return gopter.CombineGens([]gopter.Gen{
		Int64Range(0, 1),
		Int64Range(0, 0x7fe),
		Int64Range(0, 0xfffffffffffff),
	}, func(values []interface{}) interface{} {
		sign := uint64(values[0].(int64))
		exponent := uint64(values[1].(int64))
		mantissa := uint64(values[2].(int64))

		return math.Float64frombits((sign << 63) | (exponent << 52) | mantissa)
	}).WithShrinker(Float64Shrinker)
}
