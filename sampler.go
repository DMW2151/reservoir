package reservoir

// Sampler -
type Sampler interface {
	evaluateSample(int64) (int64, bool)
	reset()
}
