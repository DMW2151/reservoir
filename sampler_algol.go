package reservoir

import (
	"math"
	"math/rand"
)

// LSampler - Implements Sampler - Runs algorithm L from Devroye (1986)
// see: https://dl.acm.org/doi/pdf/10.1145/198429.198435
type LSampler struct {
	resSize int64
	rSrc    *rand.Rand
	s       int64
	w       float64
}

// reset
func (ls *LSampler) reset() {
	ls.s = 0
	ls.w = 1.0
}

// evaluateSample -
func (ls *LSampler) evaluateSample(sidx int64) (int64, bool) {

	// approve all samples w. index contained in (0, resSize)
	if sidx < ls.resSize {
		ls.s = sidx + 1
		return sidx, true
	}

	// updates condition...
	if ls.s <= sidx {
		ls.updateParameters(sidx)
		return ls.rSrc.Int63n(ls.resSize - 1), true
	}

	return 0, false
}

// updateParameters
func (ls *LSampler) updateParameters(sidx int64) {
	var wrnd, srnd = ls.rSrc.Float64(), ls.rSrc.Float64()

	ls.w *= math.Exp(math.Log(wrnd) / float64(ls.resSize))
	ls.s += 1 + int64(math.Floor(math.Log(srnd)/math.Log(1-ls.w)))

}
