package gen

import (
	"time"

	"github.com/leanovate/gopter"
)

func Time() gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		sec := genParams.NextInt64()
		usec := genParams.NextInt64()

		return gopter.NewGenResult(time.Unix(sec, usec), TimeShrinker)
	}
}

func TimeRange(from time.Time, duration time.Duration) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		v := from.Add(time.Duration(genParams.Rng.Int63n(int64(duration))))
		return gopter.NewGenResult(v, TimeShrinker)
	}
}
