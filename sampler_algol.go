package reservoir

import (
	"math"
	"math/rand"
)

// LSampler - implements Sampler - holds metadata to run algorithm-L as
// presented in Li (1994) - https://dl.acm.org/doi/pdf/10.1145/198429.198435
// see also: http://luc.devroye.org/chapter_twelve.pdf
type LSampler struct {
	resSize int64
	rSrc    *rand.Rand
	s       int64
	w       float64
}

// reset - implements Sampler's reset() method
func (ls *LSampler) reset() {
	ls.s = 0
	ls.w = 1.0
}

// evaluateSample - implements Sampler - Using the index of an observed sample (sidx), determine
// if associated data should be accepted to / rejected from reservoir
func (ls *LSampler) evaluateSample(sidx int64) (int64, bool) {

	// sample index contained in (resSize, ls.s) - do nothing,
	// upfront b.c.most common case on sufficiently large streams
	if (ls.resSize < sidx) && (sidx < ls.s) {
		return 0, false
	}

	// fill reservoir w. samples having index in [0, resSize]
	if (0 <= sidx) && (sidx < ls.resSize) {
		ls.s = sidx + 1
		return sidx, true
	}

	// sample index in (ls.s, +inf) - replace rand entry from reservoir & set
	// new params. see readme for statistical justification for modifications to
	// vals ls.w & ls.s
	ls.w *= math.Exp(math.Log(ls.rSrc.Float64()) / float64(ls.resSize))
	ls.s += 1 + int64(math.Floor(math.Log(ls.rSrc.Float64())/math.Log(1-ls.w)))

	return ls.rSrc.Int63n(ls.resSize - 1), true
}
