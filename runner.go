package gopter

import "sync"

type shouldStop func() bool

type worker func(int, shouldStop) *CheckResult

type runner struct {
	sync.RWMutex
	parameters *CheckParameters
	worker     worker
}

func (r *runner) mergeCheckResults(r1, r2 *CheckResult) *CheckResult {
	var status CheckStatus
	if r1.Status != CheckPassed && r1.Status != CheckExhausted {
		status = r1.Status
	} else if r2.Status != CheckPassed && r2.Status != CheckExhausted {
		status = r2.Status
	} else {
		if r1.Succeeded+r2.Succeeded >= r.parameters.MinSuccessfulTests &&
			float64(r1.Discarded+r2.Discarded) <= float64(r1.Succeeded+r2.Succeeded)*r.parameters.MaxDiscardRatio {
			status = CheckPassed
		} else {
			status = CheckExhausted
		}
	}
	return &CheckResult{
		Status:    status,
		Succeeded: r1.Succeeded + r2.Succeeded,
		Discarded: r1.Discarded + r2.Discarded,
	}
}

func (r *runner) runWorkers() *CheckResult {
	var stopFlag Flag
	defer stopFlag.Set()

	if r.parameters.Workers < 2 {
		return r.worker(0, stopFlag.Get)
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(r.parameters.Workers)
	results := make(chan *CheckResult, r.parameters.Workers)
	combinedResult := make(chan *CheckResult)

	go func() {
		var combined *CheckResult
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
		go func() {
			defer waitGroup.Done()
			results <- r.worker(i, stopFlag.Get)
		}()
	}
	waitGroup.Wait()
	close(results)

	return <-combinedResult
}
