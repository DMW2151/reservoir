package reservoir

import (
	"math/rand"
)

// reservoirSample - Maintains a reservoir sample
type reservoirSample[T any] struct {
	data    []T
	sampler Sampler
	seen    int64
}

// NewreservoirSample -
func NewreservoirSample[T any](s int64, seed int64) reservoirSample[T] {
	return reservoirSample[T]{
		data: make([]T, s),
		sampler: &LSampler{
			resSize: s,
			rSrc:    rand.New(rand.NewSource(seed)),
			w:       1, // see LSampler for statistical justification for init vals (w, s)
			s:       0,
		},
	}
}

// storeSample
func (rs *reservoirSample[T]) storeSample(idx int64, content T) {
	rs.data[idx] = content
}

// ReadSample -
func (rs *reservoirSample[T]) ReadSample(content T) bool {
	rs.seen += 1
	if ridx, ok := rs.sampler.evaluateSample(rs.seen); ok {
		rs.storeSample(ridx, content)
		return true
	}
	return false
}

// NumSamplesSeen
func (rs *reservoirSample[T]) NumSamplesSeen() int64 {
	return rs.seen
}

// Samples
func (rs *reservoirSample[T]) Samples() []T {
	return rs.data
}

// Reset
func (rs *reservoirSample[T]) Reset() {
	l := cap(rs.data)
	rs.data = make([]T, l)
	rs.seen = 0
	rs.sampler.reset()
}
