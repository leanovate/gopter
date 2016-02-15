package prop

import "github.com/untoldwind/gopter"

func ForAll(check Check, gens ...gopter.Gen) gopter.Prop {
	return func(genParams *gopter.GenParameters) *gopter.PropResult {
		genResults, values, ok := generatorResults(genParams, gens)
		if !ok {
			return &gopter.PropResult{
				Status: gopter.PropUndecided,
			}
		}

		result := convertResult(check(values...))
		if result.Success() {
			return result.WithArgs(noShrinkArgs(genResults, values))
		}
		return nil
	}
}
