package gopter

import (
	"math/rand"
	"time"
)

// TestParameters to run property tests
type CheckParameters struct {
	MinSuccessfulTests int
	MinSize            int
	MaxSize            int
	Rng                *rand.Rand
	Workers            int
	MaxDiscardRatio    float64
}

// DefaultTestParameters creates reasonable default Parameters for most cases
func DefaultCheckParameters() *CheckParameters {
	seed := time.Now().UnixNano()

	return &CheckParameters{
		MinSuccessfulTests: 100,
		MinSize:            0,
		MaxSize:            100,
		Rng:                rand.New(rand.NewSource(seed)),
		Workers:            1,
		MaxDiscardRatio:    5,
	}
}
