package prop

import (
	"errors"
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
	var callCheck func([]interface{}) *gopter.PropResult
	if checkType.NumOut() == 0 {
		return ErrorProp(errors.New("At least one output parameters is required"))
	} else if checkType.NumOut() > 2 {
		return ErrorProp(fmt.Errorf("No more than 2 output parameters are allowed: %d", checkType.NumOut()))
	} else if checkType.NumOut() == 2 && !checkType.Out(1).Implements(typeOfError) {
		return ErrorProp(fmt.Errorf("No 2 output has to be error: %v", checkType.Out(1).Kind()))
	} else if checkType.NumOut() == 2 {
		callCheck = func(values []interface{}) *gopter.PropResult {
			rvs := make([]reflect.Value, len(values))
			for i, value := range values {
				rvs[i] = reflect.ValueOf(value)
			}
			results := checkVal.Call(rvs)
			if results[1].IsNil() {
				return convertResult(results[0].Interface(), nil)
			}
			return convertResult(results[0].Interface(), results[1].Interface().(error))
		}
	} else {
		callCheck = func(values []interface{}) *gopter.PropResult {
			rvs := make([]reflect.Value, len(values))
			for i, value := range values {
				rvs[i] = reflect.ValueOf(value)
			}
			results := checkVal.Call(rvs)
			return convertResult(results[0].Interface(), nil)
		}
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
		if result.Success() {
			for i, genResult := range genResults {
				result = result.WithArgs(gopter.NewPropArg(genResult, 0, values[i], values[i]))
			}
		} else {
			for i, genResult := range genResults {
				result, values[i] = shrinkValue(genParams.MaxShrinkCount, genResult, values[i], result,
					func(v interface{}) *gopter.PropResult {
						shrinkedOne := make([]interface{}, len(values))
						copy(shrinkedOne, values)
						shrinkedOne[i] = v
						return callCheck(shrinkedOne)
					})
			}
		}
		return result
	})
}

func ForAll1(gen gopter.Gen, check func(v interface{}) (interface{}, error)) gopter.Prop {
	checkFunc := func(v interface{}) *gopter.PropResult {
		return convertResult(check(v))
	}
	return gopter.SaveProp(func(genParams *gopter.GenParameters) *gopter.PropResult {
		genResult := gen(genParams)
		value, ok := genResult.Retrieve()
		if !ok {
			return &gopter.PropResult{
				Status: gopter.PropUndecided,
			}
		}
		result := checkFunc(value)
		if result.Success() {
			return result.WithArgs(gopter.NewPropArg(genResult, 0, value, value))
		}

		result, _ = shrinkValue(genParams.MaxShrinkCount, genResult, value, result, checkFunc)
		return result
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
	firstFail *gopter.PropResult, check func(interface{}) *gopter.PropResult) (*gopter.PropResult, interface{}) {
	lastFail := firstFail
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

	return lastFail.WithArgs(firstFail.Args...).WithArgs(gopter.NewPropArg(genResult, shrinks, lastValue, origValue)), lastValue
}

func firstFailure(shrink gopter.Shrink, check func(interface{}) *gopter.PropResult) (*gopter.PropResult, interface{}) {
	value, ok := shrink()
	for ok {
		result := check(value)
		if !result.Success() {
			return result, value
		}
		value, ok = shrink()
	}
	return nil, nil
}
