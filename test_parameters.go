package gopter

import (
	"math/rand"
	"time"
)

// TestParameters to run property tests
type TestParameters struct {
	MinSuccessfulTests int
	MinSize            int
	MaxSize            int
	MaxShrinkCount     int
	Rng                *rand.Rand
	Workers            int
	MaxDiscardRatio    float64
}

// DefaultTestParameters creates reasonable default Parameters for most cases
func DefaultTestParameters() *TestParameters {
	seed := time.Now().UnixNano()

	return &TestParameters{
		MinSuccessfulTests: 100,
		MinSize:            0,
		MaxSize:            100,
		MaxShrinkCount:     1000,
		Rng:                rand.New(rand.NewSource(seed)),
		Workers:            1,
		MaxDiscardRatio:    5,
	}
}
