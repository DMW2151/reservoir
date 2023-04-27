package reservoir

import (
	"math/rand"
	"testing"
)

var (
	benchmarkSeed           = int64(2151)
	benchmarkTinyreservoir  = int64(16)
	benchmarkSmallreservoir = int64(256)
	benchmarkLargereservoir = int64(16 * 1024)

	benchmarkScenarios = []struct {
		name             string
		reservoirSize    int64
		seed             int64
		samplesToProcces int
	}{
		{
			name:             "1M_tiny_reservoir",
			reservoirSize:    benchmarkTinyreservoir,
			seed:             benchmarkSeed,
			samplesToProcces: 1024,
		},
		{
			name:             "1M_small_reservoir",
			reservoirSize:    benchmarkSmallreservoir,
			seed:             benchmarkSeed,
			samplesToProcces: 1024,
		},
		{
			name:             "1M_large_reservoir",
			reservoirSize:    benchmarkLargereservoir,
			seed:             benchmarkSeed,
			samplesToProcces: 1024 * 1024,
		},
		{
			name:             "8M_tiny_reservoir",
			reservoirSize:    benchmarkTinyreservoir,
			seed:             benchmarkSeed,
			samplesToProcces: 1024 * 1024 * 8,
		},
		{
			name:             "8M_small_reservoir",
			reservoirSize:    benchmarkSmallreservoir,
			seed:             benchmarkSeed,
			samplesToProcces: 1024 * 1024 * 8,
		},
		{
			name:             "8M_large_reservoir",
			reservoirSize:    benchmarkLargereservoir,
			seed:             benchmarkSeed,
			samplesToProcces: 1024 * 1024 * 8,
		},
	}
)

// Benchmark_EvaluateSampler_AlgorithmL - Benchmark algoL implementation
func Benchmark_EvaluateSampler_AlgorithmL(b *testing.B) {
	for _, bm := range benchmarkScenarios {
		b.Run(bm.name, func(b *testing.B) {

			for n := 0; n < b.N; n++ {

				b.StopTimer()
				ls := LSampler{
					resSize: bm.reservoirSize,
					rSrc:    rand.New(rand.NewSource(bm.seed)),
					w:       1, // see LSampler for statistical justification for init vals (w, s)
					s:       bm.reservoirSize,
				}

				b.StartTimer()

				for i := 0; i < bm.samplesToProcces; i++ {
					ls.evaluateSample(int64(i))
				}
			}
		})
	}
}
