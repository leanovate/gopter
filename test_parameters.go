package gopter

import (
	"math/rand"
	"time"
)

// TestParameters to run property tests
type TestParameters struct {
	MinSuccessfulTests int
	// MinSize is an (inclusive) lower limit on the size of the parameters
	MinSize int
	// MaxSize is an (exclusive) upper limit on the size of the parameters
	MaxSize         int
	MaxShrinkCount  int
	Seed            int64
	Rng             *rand.Rand
	Workers         int
	MaxDiscardRatio float64
}

// DefaultTestParameters creates reasonable default Parameters for most cases
func DefaultTestParameters() *TestParameters {
	seed := time.Now().UnixNano()

	return &TestParameters{
		MinSuccessfulTests: 100,
		MinSize:            0,
		MaxSize:            100,
		MaxShrinkCount:     1000,
		Seed:               seed,
		Rng:                rand.New(rand.NewSource(seed)),
		Workers:            1,
		MaxDiscardRatio:    5,
	}
}
