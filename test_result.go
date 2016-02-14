package gopter

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
}

// Passed checks if the check has passed
func (r *TestResult) Passed() bool {
	return r.Status == TestPassed || r.Status == TestProved
}
