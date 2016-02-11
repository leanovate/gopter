package gopter

import (
	"math/rand"
	"time"
)

type Parameters struct {
	MinSuccessfulTests int
	MinSize            int
	MaxSize            int
	Rng                *rand.Rand
}

func DefaultParameters() *Parameters {
	seed := time.Now().UnixNano()

	return &Parameters{
		MinSuccessfulTests: 100,
		MinSize:            0,
		MaxSize:            100,
		Rng:                rand.New(rand.NewSource(seed)),
	}
}
