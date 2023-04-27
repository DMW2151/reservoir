package reservoir

// Sampler - For Sampling Algos...
type Sampler interface {
	evaluateSample(int64) (int64, bool)
	reset()
}
