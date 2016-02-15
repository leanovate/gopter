package gen

import "github.com/untoldwind/gopter"

func OneConstOf(first interface{}, other ...interface{}) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		idx := genParams.Rng.Intn(len(other) + 1)
		if idx == 0 {
			gopter.NewGenResult(first, gopter.NoShrinker)
		}
		return gopter.NewGenResult(other[idx-1], gopter.NoShrinker)
	}
}

func OneGenOf(first gopter.Gen, other ...gopter.Gen) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		idx := genParams.Rng.Intn(len(other) + 1)
		if idx == 0 {
			return first(genParams)
		}
		return other[idx+1](genParams)
	}
}
