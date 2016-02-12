package prop

import (
	"fmt"

	"github.com/untoldwind/gopter"
)

// CheckCondition ecapsulates a condition by which a property can be chacked
// This is what testers usually have to implement
type CheckCondition func(...interface{}) (interface{}, error)

func NewCheck1(f func(interface{}) (interface{}, error)) CheckCondition {
	return func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Expected 1 argument, got %d", len(args))
		}
		return f(args[0])
	}
}

func convertResult(result interface{}, err error) *gopter.PropResult {
	if err != nil {
		return &gopter.PropResult{
			Status: gopter.PropError,
			Error:  err,
		}
	}
	switch result.(type) {
	case bool:
		if result.(bool) {
			return &gopter.PropResult{Status: gopter.PropTrue}
		}
		return &gopter.PropResult{Status: gopter.PropFalse}
	case *gopter.PropResult:
		return result.(*gopter.PropResult)
	}
	return &gopter.PropResult{
		Status: gopter.PropError,
		Error:  fmt.Errorf("Invalid check result: %#v", result),
	}
}
