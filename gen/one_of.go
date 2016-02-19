package gen

import "github.com/leanovate/gopter"

// OneConstOf generate one of a list of constant values
func OneConstOf(first interface{}, other ...interface{}) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		idx := genParams.Rng.Intn(len(other) + 1)
		if idx == 0 {
			return gopter.NewGenResult(first, gopter.NoShrinker)
		}
		return gopter.NewGenResult(other[idx-1], gopter.NoShrinker)
	}
}

// OneGenOf generate one value from a a list of generators
func OneGenOf(first gopter.Gen, other ...gopter.Gen) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		idx := genParams.Rng.Intn(len(other) + 1)
		if idx == 0 {
			return first(genParams)
		}
		return other[idx+1](genParams)
	}
}
