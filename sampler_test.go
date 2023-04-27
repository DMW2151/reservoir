package reservoir

import (
	"math/rand"
	"testing"
)

var (
	TestSeed           = int64(2151)
	TestMicroreservoir = int64(4)
)

var testScenarios = []struct {
	name             string
	reservoirSize    int64
	seed             int64
	samplesToProcces int
}{
	{
		name:             "1M_test_approximate_equal_sampling",
		reservoirSize:    TestMicroreservoir,
		seed:             TestSeed,
		samplesToProcces: 100,
	},
}

func Test_Seeded_Samples_AlgoR(t *testing.T) {

	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {

			var updateFreqArr = make([]int64, test.reservoirSize)

			rs := RSampler{
				resSize: test.reservoirSize,
				rSrc:    rand.New(rand.NewSource(test.seed)),
			}

			for i := 0; i < test.samplesToProcces; i++ {
				if replaceIndex, replace := rs.evaluateSample(int64(i)); replace {
					updateFreqArr[replaceIndex] += 1
				}
			}
			t.Log(updateFreqArr)
		})
	}
}

func Test_Seeded_Samples_AlgoL(t *testing.T) {

	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {

			var updateFreqArr = make([]int64, test.reservoirSize)

			ls := LSampler{
				resSize: test.reservoirSize,
				rSrc:    rand.New(rand.NewSource(test.seed)),
				w:       1, // see LSampler for statistical justification for init vals (w, s)
				s:       0,
			}

			for i := 0; i < test.samplesToProcces; i++ {
				if replaceIndex, replace := ls.evaluateSample(int64(i)); replace {
					updateFreqArr[replaceIndex] += 1
				}
			}
			t.Log(updateFreqArr)
		})
	}
}
