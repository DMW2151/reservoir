package reservoir

import (
	"math/rand"
)

// RSampler - Implements Sampler - Run algorithm R from Vitter (1985)
// see: https://www.cs.umd.edu/~samir/498/vitter.pdf
type RSampler struct {
	resSize int64
	rSrc    *rand.Rand
}

// reset -
func (rs *RSampler) reset() {}

// evaluateSample -
func (rs *RSampler) evaluateSample(sidx int64) (int64, bool) {

	// approve all samples w. index in (0, resSize)
	if sidx < rs.resSize {
		return sidx, true
	}

	// generate a uniformly distributed integer, k ~ U(1, `sidx`), accept into
	// reservoir when k ⊆ (1, `resSize`).
	if k := rs.rSrc.Int63n(sidx - 1); k < rs.resSize {
		return k, true
	}
	return 0, false
}

func (rs RSampler) size() int64 {
	return rs.resSize
}
