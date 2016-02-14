package gopter

import (
	"sync/atomic"
	"testing"
)

func TestPropPassed(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropTrue,
		}
	})

	parameters := DefaultTestParameters()
	result := prop.Check(parameters)

	if result.Status != TestPassed || result.Succeeded != parameters.MinSuccessfulTests {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != int64(parameters.MinSuccessfulTests) {
		t.Errorf("Invalid number of calls: %d", called)
	}
}

func TestPropPassedMulti(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropTrue,
		}
	})

	parameters := DefaultTestParameters()
	parameters.Workers = 10
	result := prop.Check(parameters)

	if result.Status != TestPassed || result.Succeeded != parameters.MinSuccessfulTests {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != int64(parameters.MinSuccessfulTests) {
		t.Errorf("Invalid number of calls: %d", called)
	}
}
