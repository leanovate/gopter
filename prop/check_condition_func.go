package prop

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/leanovate/gopter"
)

type typedValue struct {
	value       interface{}
	reflectType reflect.Type
}

func (t typedValue) ValueOf() reflect.Value {
	if t.value == nil {
		return reflect.Zero(t.reflectType)
	}
	return reflect.ValueOf(t.value)
}

func checkConditionFunc(check interface{}, numArgs int) (func([]typedValue) *gopter.PropResult, error) {
	checkVal := reflect.ValueOf(check)
	checkType := checkVal.Type()

	if checkType.Kind() != reflect.Func {
		return nil, fmt.Errorf("First param of ForrAll has to be a func: %v", checkVal.Kind())
	}
	if checkType.NumIn() != numArgs {
		return nil, fmt.Errorf("Number of parameters does not match number of generators: %d != %d", checkType.NumIn(), numArgs)
	}
	if checkType.NumOut() == 0 {
		return nil, errors.New("At least one output parameters is required")
	} else if checkType.NumOut() > 2 {
		return nil, fmt.Errorf("No more than 2 output parameters are allowed: %d", checkType.NumOut())
	} else if checkType.NumOut() == 2 && !checkType.Out(1).Implements(typeOfError) {
		return nil, fmt.Errorf("No 2 output has to be error: %v", checkType.Out(1).Kind())
	} else if checkType.NumOut() == 2 {
		return func(values []typedValue) *gopter.PropResult {
			rvs := make([]reflect.Value, len(values))
			for i, value := range values {
				rvs[i] = value.ValueOf()
			}
			results := checkVal.Call(rvs)
			if results[1].IsNil() {
				return convertResult(results[0].Interface(), nil)
			}
			return convertResult(results[0].Interface(), results[1].Interface().(error))
		}, nil
	}
	return func(values []typedValue) *gopter.PropResult {
		rvs := make([]reflect.Value, len(values))
		for i, value := range values {
			rvs[i] = value.ValueOf()
		}
		results := checkVal.Call(rvs)
		return convertResult(results[0].Interface(), nil)
	}, nil
}
