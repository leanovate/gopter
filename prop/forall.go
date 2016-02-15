package prop

import "github.com/leanovate/gopter"

func ForAll1(gen gopter.Gen, check func(interface{}) (interface{}, error)) gopter.Prop {
	return gopter.SaveProp(func(genParams *gopter.GenParameters) *gopter.PropResult {
		genResult := gen(genParams)
		value, ok := genResult.Retrieve()
		if !ok {
			return &gopter.PropResult{
				Status: gopter.PropUndecided,
			}
		}
		result := convertResult(check(value))
		if result.Success() {
			return result.WithArg(gopter.NewPropArg(genResult, 0, value, value))
		}

		return shrinkValue(genResult, value, result, check)
	})
}

func shrinkValue(genResult *gopter.GenResult, origValue interface{},
	lastFail *gopter.PropResult, check func(interface{}) (interface{}, error)) *gopter.PropResult {
	lastValue := origValue

	shrinks := 0
	shrink := genResult.Shrinker(lastValue).Filter(genResult.Sieve)
	nextResult, nextValue := firstFailure(shrink, check)
	for nextResult != nil {
		shrinks++
		lastValue = nextValue
		lastFail = nextResult

		shrink = genResult.Shrinker(lastValue).Filter(genResult.Sieve)
		nextResult, nextValue = firstFailure(shrink, check)
	}

	return lastFail.WithArg(gopter.NewPropArg(genResult, shrinks, lastValue, origValue))
}

func firstFailure(shrink gopter.Shrink, check func(interface{}) (interface{}, error)) (*gopter.PropResult, interface{}) {
	value, ok := shrink()
	for ok {
		result := convertResult(check(value))
		if !result.Success() {
			return result, value
		}
		value, ok = shrink()
	}
	return nil, nil
}
