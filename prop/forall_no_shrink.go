package prop

import "github.com/leanovate/gopter"

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
		return convertResult(check(value)).WithArg(gopter.NewPropArg(genResult, 0, value, value))
	})
}

func ForAllNoShrink2(gen1, gen2 gopter.Gen, check func(v1, v2 interface{}) (interface{}, error)) gopter.Prop {
	return ForAllNoShrink1(gen1, func(v1 interface{}) (interface{}, error) {
		return ForAllNoShrink1(gen2, func(v2 interface{}) (interface{}, error) {
			return check(v1, v2)
		}), nil
	})
}
