package gopter

import (
	"math/rand"
	"time"
)

// Parameters to run property tests
type Parameters struct {
	MinSuccessfulTests int
	MinSize            int
	MaxSize            int
	Rng                *rand.Rand
	Workers            int
	MaxDiscardRatio    float64
}

// DefaultParameters creates reasonable default Parameters for most cases
func DefaultParameters() *Parameters {
	seed := time.Now().UnixNano()

	return &Parameters{
		MinSuccessfulTests: 100,
		MinSize:            0,
		MaxSize:            100,
		Rng:                rand.New(rand.NewSource(seed)),
		Workers:            1,
		MaxDiscardRatio:    5,
	}
}
