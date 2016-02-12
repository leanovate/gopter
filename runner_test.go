package gopter

import "testing"

func TestRunnerSingleWorker(t *testing.T) {
	parameters := DefaultTestParameters()
	testRunner := &runner{
		parameters: parameters,
		worker: func(num int, shouldStop shouldStop) *TestResult {
			return &TestResult{
				Status:    Passed,
				Succeeded: 1,
				Discarded: 0,
			}
		},
	}

	result := testRunner.runWorkers()

	if result.Status != Passed ||
		result.Succeeded != 1 ||
		result.Discarded != 0 {
		t.Errorf("Invalid result: %#v", result)
	}
}

func TestRunnerParallelWorkers(t *testing.T) {
	parameters := DefaultTestParameters()
	parameters.Workers = 50
	testRunner := &runner{
		parameters: parameters,
		worker: func(num int, shouldStop shouldStop) *TestResult {
			return &TestResult{
				Status:    Passed,
				Succeeded: 10,
				Discarded: 1,
			}
		},
	}

	result := testRunner.runWorkers()

	if result.Status != Passed ||
		result.Succeeded != 500 ||
		result.Discarded != 50 {
		t.Errorf("Invalid result: %#v", result)
	}
}
