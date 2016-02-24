package gopter

import (
	"math/rand"
	"time"
)

type GenParameters struct {
	Size           int
	MaxShrinkCount int
	Rng            *rand.Rand
}

func (p *GenParameters) WithSize(size int) *GenParameters {
	newParameters := *p
	newParameters.Size = size
	return &newParameters
}

func (p *GenParameters) NextBool() bool {
	return p.Rng.Int63()&1 == 0
}

func (p *GenParameters) NextInt64() int64 {
	v := p.Rng.Int63()
	if p.NextBool() {
		return -v
	}
	return v
}

func (p *GenParameters) NextUint64() uint64 {
	first := uint64(p.Rng.Int63())
	second := uint64(p.Rng.Int63())

	return (first << 1) ^ second
}

func DefaultGenParameters() *GenParameters {
	seed := time.Now().UnixNano()

	return &GenParameters{
		Size:           100,
		MaxShrinkCount: 1000,
		Rng:            rand.New(rand.NewSource(seed)),
	}
}
