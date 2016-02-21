package gopter

import (
	"fmt"
	"time"
)

type testStatus int

const (
	TestPassed testStatus = iota
	TestProved
	TestFailed
	TestExhausted
	TestError
)

func (s testStatus) String() string {
	switch s {
	case TestPassed:
		return "PASSED"
	case TestProved:
		return "PROVED"
	case TestFailed:
		return "FAILED"
	case TestExhausted:
		return "EXHAUSTED"
	case TestError:
		return "ERROR"
	}
	return ""
}

type TestResult struct {
	Status    testStatus
	Succeeded int
	Discarded int
	Error     error
	Args      PropArgs
	Time      time.Duration
}

// Passed checks if the check has passed
func (r *TestResult) Passed() bool {
	return r.Status == TestPassed || r.Status == TestProved
}

func (r *TestResult) Report() string {
	status := ""
	switch r.Status {
	case TestProved:
		status = "OK, proved property.\n" + r.Args.Report()
	case TestPassed:
		status = fmt.Sprintf("OK, passed %d tests.", r.Succeeded)
	case TestFailed:
		status = fmt.Sprintf("Falsified after %d passed tests.\n", r.Succeeded) + r.Args.Report()
	case TestExhausted:
		status = fmt.Sprintf("Gave up after only %d passed tests. %d tests were discarded.", r.Succeeded, r.Discarded)
	case TestError:
		status = fmt.Sprintf("Error on property evaluation: %s\n", r.Succeeded, r.Error.Error()) + r.Args.Report()
	}

	return concatLines(status, fmt.Sprintf("Elapsed time: %s", r.Time.String()))
}
