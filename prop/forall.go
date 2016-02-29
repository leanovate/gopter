package prop

import (
	"fmt"
	"reflect"

	"github.com/leanovate/gopter"
)

var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

func ForAll(check interface{}, gens ...gopter.Gen) gopter.Prop {
	checkVal := reflect.ValueOf(check)
	checkType := checkVal.Type()
	if checkType.Kind() != reflect.Func {
		return ErrorProp(fmt.Errorf("First param of ForrAll has to be a func: %v", checkVal.Kind()))
	}
	if checkType.NumIn() != len(gens) {
		return ErrorProp(fmt.Errorf("Number of parameters does not match number of generators: %d != %d", checkType.NumIn(), len(gens)))
	}
	if checkType.NumOut() > 2 {
		return ErrorProp(fmt.Errorf("No more than 2 output parameters are allowed: %d", checkType.NumOut()))
	} else if checkType.NumOut() == 2 && !checkType.Out(1).Implements(typeOfError) {
		return ErrorProp(fmt.Errorf("No 2 output has to be error: %v", checkType.Out(1).Kind()))
	}
	gen := gopter.CombineGens(gens...)

	return gopter.SaveProp(func(genParams *gopter.GenParameters) *gopter.PropResult {
		genResult := gen(genParams)
		value, ok := genResult.Retrieve()
		if !ok {
			return &gopter.PropResult{
				Status: gopter.PropUndecided,
			}
		}
		values := value.([]interface{})
		rvs := make([]reflect.Value, len(values))
		for i, value := range values {
			rvs[i] = reflect.ValueOf(value)
		}
		results := checkVal.Call(rvs)
		var result *gopter.PropResult
		if len(results) == 2 {
			result = convertResult(results[0].Interface(), results[1].Interface().(error))
		} else {
			result = convertResult(results[0].Interface(), nil)
		}
		args := make([]*gopter.PropArg, len(rvs))
		for i, value := range values {
			args[i] = gopter.NewPropArg(genResult, 0, value, value)
		}
		return result.WithArgs(args...)
	})
}

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
			return result.WithArgs(gopter.NewPropArg(genResult, 0, value, value))
		}

		return shrinkValue(genParams.MaxShrinkCount, genResult, value, result, check)
	})
}

func ForAll2(gen1, gen2 gopter.Gen, check func(v1, v2 interface{}) (interface{}, error)) gopter.Prop {
	return ForAll1(gen1, func(v1 interface{}) (interface{}, error) {
		return ForAll1(gen2, func(v2 interface{}) (interface{}, error) {
			return check(v1, v2)
		}), nil
	})
}

func shrinkValue(maxShrinkCount int, genResult *gopter.GenResult, origValue interface{},
	lastFail *gopter.PropResult, check func(interface{}) (interface{}, error)) *gopter.PropResult {
	lastValue := origValue

	shrinks := 0
	shrink := genResult.Shrinker(lastValue).Filter(genResult.Sieve)
	nextResult, nextValue := firstFailure(shrink, check)
	for nextResult != nil && shrinks < maxShrinkCount {
		shrinks++
		lastValue = nextValue
		lastFail = nextResult

		shrink = genResult.Shrinker(lastValue).Filter(genResult.Sieve)
		nextResult, nextValue = firstFailure(shrink, check)
	}

	return lastFail.WithArgs(gopter.NewPropArg(genResult, shrinks, lastValue, origValue))
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
