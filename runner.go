package gopter

import "sync"

type TestStatus int

const (
	Passed TestStatus = iota
	Proved
	Failed
	Exhausted
	PropError
)

type TestResult struct {
	Status    TestStatus
	Succeeded int
	Discarded int
}

type shouldStop func() bool

type worker func(int, shouldStop) *TestResult

type runner struct {
	sync.RWMutex
	parameters *TestParameters
	worker     worker
}

func (r *runner) mergeTestResults(r1, r2 *TestResult) *TestResult {
	var status TestStatus
	if r1.Status != Passed && r1.Status != Exhausted {
		status = r1.Status
	} else if r2.Status != Passed && r2.Status != Exhausted {
		status = r2.Status
	} else {
		if r1.Succeeded+r2.Succeeded >= r.parameters.MinSuccessfulTests &&
			float64(r1.Discarded+r2.Discarded) <= float64(r1.Succeeded+r2.Succeeded)*r.parameters.MaxDiscardRatio {
			status = Passed
		} else {
			status = Exhausted
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

	if r.parameters.Workers < 2 {
		return r.worker(0, stopFlag.Get)
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
				combined = r.mergeTestResults(combined, result)
			}
		}
		combinedResult <- combined
	}()
	for i := 0; i < r.parameters.Workers; i++ {
		go func() {
			defer waitGroup.Done()
			results <- r.worker(i, stopFlag.Get)
		}()
	}
	waitGroup.Wait()
	close(results)

	return <-combinedResult
}
