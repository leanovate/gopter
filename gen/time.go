package gen

import (
	"time"

	"github.com/leanovate/gopter"
)

// Time generates an arbitrary time.Time
func Time() gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		sec := genParams.NextInt64()
		usec := genParams.NextInt64()

		return gopter.NewGenResult(time.Unix(sec, usec), TimeShrinker)
	}
}

// TimeRange generates an arbitrary time.Time with a range
// from defines the start of the time range
// duration defines the overall duration of the time range
func TimeRange(from time.Time, duration time.Duration) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		v := from.Add(time.Duration(genParams.Rng.Int63n(int64(duration))))
		return gopter.NewGenResult(v, TimeShrinker)
	}
}
