package prop

import (
	"strings"

	"github.com/untoldwind/gopter"
)

func noShrinkArgs(genResults []*gopter.GenResult, values []interface{}) []gopter.PropArg {
	result := make([]gopter.PropArg, len(genResults))
	for i, genResult := range genResults {
		result[i] = gopter.PropArg{
			Label:   strings.Join(genResult.Labels, ", "),
			Arg:     values[i],
			OrigArg: values[i],
			Shrinks: 0,
		}
	}
	return result
}

// ForAllNoShrink creates a property that requires the check condition to be true for all values
func ForAllNoShrink(checkCondition CheckCondition, gens ...gopter.Gen) gopter.Prop {
	return func(genParams *gopter.GenParameters) *gopter.PropResult {
		genResults := make([]*gopter.GenResult, len(gens))
		values := make([]interface{}, len(gens))
		var ok bool
		for i, gen := range gens {
			genResults[i] = gen(genParams)
			values[i], ok = genResults[i].Retrieve()
			if !ok {
				return &gopter.PropResult{
					Status: gopter.PropUndecided,
				}
			}
		}
		return convertResult(checkCondition(values...)).WithArgs(noShrinkArgs(genResults, values))
	}
}
