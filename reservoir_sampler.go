package reservoir

import (
	"math/rand"
)

// ReservoirSample - Samples from uniformly from a stream of unknown length
type ReservoirSample[T any] struct {
	reservoir []T
	sampler   Sampler
	seen      int64
}

// NewReservoirSample - creates new ReservoirSample using Algo L
func NewReservoirSample[T any](s int64, seed int64) ReservoirSample[T] {
	return ReservoirSample[T]{
		reservoir: make([]T, s),
		sampler: &LSampler{
			resSize: s,
			rSrc:    rand.New(rand.NewSource(seed)),
			w:       1, // see readme for statistical justification for vals (w, s)
			s:       0,
		},
	}
}

// ReadSample - Consider new data for reservoir. Will either drop data or
// write to reservoir using logic in ReservoirSample.sampler
func (rs *ReservoirSample[T]) ReadSample(data T) bool {
	rs.seen += 1
	if ridx, ok := rs.sampler.evaluateSample(rs.seen); ok {
		rs.storeSample(ridx, data)
		return true
	}
	return false
}

// NumSamplesSeen - return number of samples read since init or last
// call to Reset()
func (rs *ReservoirSample[T]) NumSamplesSeen() int64 {
	return rs.seen
}

// Samples - return all current reservoir samples
func (rs *ReservoirSample[T]) Samples() []T {
	return rs.reservoir
}

// Reset - reset all statistics on the reservoirSample
func (rs *ReservoirSample[T]) Reset() {
	l := cap(rs.reservoir)
	rs.reservoir = make([]T, l)
	rs.seen = 0
	rs.sampler.reset()
}

// storeSample - sets some data into the reservoirSample's current reservoir
func (rs *ReservoirSample[T]) storeSample(idx int64, data T) {
	rs.reservoir[idx] = data
}
