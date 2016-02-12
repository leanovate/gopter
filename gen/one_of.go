package gen

import "github.com/untoldwind/gopter"

func OneConstOf(values ...interface{}) gopter.Gen {
	if len(values) == 0 {
		return Fail
	}
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		idx := genParams.Rng.Intn(len(values))
		return gopter.NewGenResult(values[idx], gopter.NoShrinker)
	}
}

func OneGenOf(generators ...gopter.Gen) gopter.Gen {
	if len(generators) == 0 {
		return Fail
	}
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		idx := genParams.Rng.Intn(len(generators))
		return generators[idx](genParams)
	}
}
