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

func (p *GenParameters) NextUint64() uint64 {
	first := uint64(p.Rng.Int63())
	second := uint64(p.Rng.Int63())

	return (first << 1) ^ second
}

func DefaultGenParameters() *GenParameters {
	seed := time.Now().UnixNano()

	return &GenParameters{
		Size: 100,
		Rng:  rand.New(rand.NewSource(seed)),
	}
}
