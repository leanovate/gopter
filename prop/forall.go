package prop

import "github.com/leanovate/gopter"

func ForAll1(gen gopter.Gen, check func(v interface{}) (interface{}, error)) gopter.Prop {
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

func ForAll2(gen1, gen2 gopter.Gen, check func(v1, v2 interface{}) (interface{}, error)) gopter.Prop {
	return ForAll1(gen1, func(v1 interface{}) (interface{}, error) {
		return ForAll1(gen2, func(v2 interface{}) (interface{}, error) {
			return check(v1, v2)
		}), nil
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
