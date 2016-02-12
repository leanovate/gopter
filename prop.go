package gopter

import "math"

type Prop func(*GenParameters) Result

func (prop Prop) check(parameters *TestParameters) *TestResult {
	iterations := math.Ceil(float64(parameters.MinSuccessfulTests) / float64(parameters.Workers))
	sizeStep := float64(parameters.MaxSize-parameters.MinSize) / (iterations * float64(parameters.Workers))
	genParameters := GenParameters{
		Rng: parameters.Rng,
	}
	runner := &runner{
		parameters: parameters,
		worker: func(workerIdx int, shouldStop shouldStop) *TestResult {
			var n int
			var d int

			isExhaused := func() bool {
				return n+d > parameters.MinSuccessfulTests &&
					float64(1+parameters.Workers*n)*parameters.MaxDiscardRatio < float64(d)
			}

			for !shouldStop() {
				size := float64(parameters.MinSize) + (sizeStep * float64(workerIdx+(parameters.Workers*(n+d))))
				propResult := prop(genParameters.WithSize(int(size)))

				switch propResult.Status {
				case Undecided:
					d++
					if isExhaused() {
						return &TestResult{
							Status:    Exhausted,
							Succeeded: n,
							Discarded: d,
						}
					}
				case True:
					n++
				case Proof:
					n++
					return &TestResult{
						Status:    Proved,
						Succeeded: n,
						Discarded: d,
					}
				case False:
					return &TestResult{
						Status:    Failed,
						Succeeded: n,
						Discarded: d,
					}
				case Error:
					return &TestResult{
						Status:    PropError,
						Succeeded: n,
						Discarded: d,
					}
				}
			}

			if isExhaused() {
				return &TestResult{
					Status:    Exhausted,
					Succeeded: n,
					Discarded: d,
				}
			}
			return &TestResult{
				Status:    Passed,
				Succeeded: n,
				Discarded: d,
			}
		},
	}

	return runner.runWorkers()
}
