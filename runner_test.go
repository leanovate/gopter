package gopter

import "testing"

func TestRunnerSingleWorker(t *testing.T) {
	parameters := DefaultCheckParameters()
	testRunner := &runner{
		parameters: parameters,
		worker: func(num int, shouldStop shouldStop) *CheckResult {
			return &CheckResult{
				Status:    CheckPassed,
				Succeeded: 1,
				Discarded: 0,
			}
		},
	}

	result := testRunner.runWorkers()

	if result.Status != CheckPassed ||
		result.Succeeded != 1 ||
		result.Discarded != 0 {
		t.Errorf("Invalid result: %#v", result)
	}
}

func TestRunnerParallelWorkers(t *testing.T) {
	parameters := DefaultCheckParameters()
	parameters.Workers = 50
	testRunner := &runner{
		parameters: parameters,
		worker: func(num int, shouldStop shouldStop) *CheckResult {
			return &CheckResult{
				Status:    CheckPassed,
				Succeeded: 10,
				Discarded: 1,
			}
		},
	}

	result := testRunner.runWorkers()

	if result.Status != CheckPassed ||
		result.Succeeded != 500 ||
		result.Discarded != 50 {
		t.Errorf("Invalid result: %#v", result)
	}
}
