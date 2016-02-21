package gopter

import (
	"sync"
	"time"
)

type shouldStop func() bool

type worker func(int, shouldStop) *TestResult

type runner struct {
	sync.RWMutex
	parameters *TestParameters
	worker     worker
}

func (r *runner) mergeCheckResults(r1, r2 *TestResult) *TestResult {
	var status testStatus
	if r1.Status != TestPassed && r1.Status != TestExhausted {
		status = r1.Status
	} else if r2.Status != TestPassed && r2.Status != TestExhausted {
		status = r2.Status
	} else {
		if r1.Succeeded+r2.Succeeded >= r.parameters.MinSuccessfulTests &&
			float64(r1.Discarded+r2.Discarded) <= float64(r1.Succeeded+r2.Succeeded)*r.parameters.MaxDiscardRatio {
			status = TestPassed
		} else {
			status = TestExhausted
		}
	}
	return &TestResult{
		Status:    status,
		Succeeded: r1.Succeeded + r2.Succeeded,
		Discarded: r1.Discarded + r2.Discarded,
	}
}

func (r *runner) runWorkers() *TestResult {
	var stopFlag Flag
	defer stopFlag.Set()

	start := time.Now()
	if r.parameters.Workers < 2 {
		result := r.worker(0, stopFlag.Get)
		result.Time = time.Since(start)
		return result
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(r.parameters.Workers)
	results := make(chan *TestResult, r.parameters.Workers)
	combinedResult := make(chan *TestResult)

	go func() {
		var combined *TestResult
		for result := range results {
			if combined == nil {
				combined = result
			} else {
				combined = r.mergeCheckResults(combined, result)
			}
		}
		combinedResult <- combined
	}()
	for i := 0; i < r.parameters.Workers; i++ {
		go func(workerIdx int) {
			defer waitGroup.Done()
			results <- r.worker(workerIdx, stopFlag.Get)
		}(i)
	}
	waitGroup.Wait()
	close(results)

	result := <-combinedResult
	result.Time = time.Since(start)
	return result
}
