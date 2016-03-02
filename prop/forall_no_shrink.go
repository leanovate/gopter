package prop

import "github.com/leanovate/gopter"

/*
ForAllNoShrink creates a property that requires the check condition to be true for all values.
As the name suggests the generated values will not be shrinked if the condition falsiies.

"condition" has to be a function with the same number of parameters as the provided
generators "gens". The function may return a simple bool, a *PropResult, a boolean with error or
a *PropResult with error.
*/
func ForAllNoShrink(condition interface{}, gens ...gopter.Gen) gopter.Prop {
	callCheck, err := checkConditionFunc(condition, len(gens))
	if err != nil {
		return ErrorProp(err)
	}

	return gopter.SaveProp(func(genParams *gopter.GenParameters) *gopter.PropResult {
		genResults := make([]*gopter.GenResult, len(gens))
		values := make([]interface{}, len(gens))
		var ok bool
		for i, gen := range gens {
			result := gen(genParams)
			genResults[i] = result
			values[i], ok = result.Retrieve()
			if !ok {
				return &gopter.PropResult{
					Status: gopter.PropUndecided,
				}
			}
		}
		result := callCheck(values)
		for i, genResult := range genResults {
			result = result.WithArgs(gopter.NewPropArg(genResult, 0, values[i], values[i]))
		}
		return result
	})
}

// ForAllNoShrink1 creates a property that requires the check condition to be true for all values
// As the name suggests the generated values will not be shrinked if the condition falsiies
func ForAllNoShrink1(gen gopter.Gen, check func(interface{}) (interface{}, error)) gopter.Prop {
	return gopter.SaveProp(func(genParams *gopter.GenParameters) *gopter.PropResult {
		genResult := gen(genParams)
		value, ok := genResult.Retrieve()
		if !ok {
			return &gopter.PropResult{
				Status: gopter.PropUndecided,
			}
		}
		return convertResult(check(value)).WithArgs(gopter.NewPropArg(genResult, 0, value, value))
	})
}

func ForAllNoShrink2(gen1, gen2 gopter.Gen, check func(v1, v2 interface{}) (interface{}, error)) gopter.Prop {
	return ForAllNoShrink1(gen1, func(v1 interface{}) (interface{}, error) {
		return ForAllNoShrink1(gen2, func(v2 interface{}) (interface{}, error) {
			return check(v1, v2)
		}), nil
	})
}
