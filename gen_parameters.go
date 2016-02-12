package gopter

import (
	"math/rand"
	"time"
)

type GenParameters struct {
	Size int
	Rng  *rand.Rand
}

func (p *GenParameters) WithSize(size int) *GenParameters {
	newParameters := *p
	newParameters.Size = size
	return &newParameters
}

var DefaultGenParameters = func() *GenParameters {
	seed := time.Now().UnixNano()

	return &GenParameters{
		Size: 100,
		Rng:  rand.New(rand.NewSource(seed)),
	}
}()
