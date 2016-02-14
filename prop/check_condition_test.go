package prop

import (
	"errors"
	"testing"

	"github.com/untoldwind/gopter"
)

func TestCheckCondition(t *testing.T) {
	trueCondition := Check(func(...interface{}) (interface{}, error) {
		return true, nil
	})

	trueResult := convertResult(trueCondition())
	if trueResult.Status != gopter.PropTrue || trueResult.Error != nil {
		t.Errorf("Invalid true result: %#v", trueResult)
	}

	falseCondition := Check(func(...interface{}) (interface{}, error) {
		return false, nil
	})

	falseResult := convertResult(falseCondition())
	if falseResult.Status != gopter.PropFalse || falseResult.Error != nil {
		t.Errorf("Invalid false result: %#v", falseResult)
	}

	errorCondition := Check(func(...interface{}) (interface{}, error) {
		return "Anthing", errors.New("Booom")
	})

	errorResult := convertResult(errorCondition())
	if errorResult.Status != gopter.PropError || errorResult.Error == nil || errorResult.Error.Error() != "Booom" {
		t.Errorf("Invalid error result: %#v", errorResult)
	}

	propCondition := Check(func(...interface{}) (interface{}, error) {
		return &gopter.PropResult{
			Status: gopter.PropProof,
		}, nil
	})

	propResult := convertResult(propCondition())
	if propResult.Status != gopter.PropProof || falseResult.Error != nil {
		t.Errorf("Invalid prop result: %#v", propResult)
	}
}
