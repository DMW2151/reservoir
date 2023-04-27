# Reservoir Sampling - Algo L

Implements Algorithm-L from *Reservoir-Sampling Algorithms of Time Complexity O(n(1 + log(N/n)))* ([Li, 1994](https://dl.acm.org/doi/10.1145/198429.198435))

## Motivation

In some circumstances, one may wish to construct a uniform sample of fixed size ($n$) over a stream of events (e.g. traces, requests) of unknown size ($N$). An accurate solution would ensure that each element in the stream has equal probability of being included in the sample.

One solution (Algorithm R) to this problem is to include the first $n$ events of the stream in the sample and then include the i-th event with probability $\frac{n}{i}$. On large streams, this algorithm becomes less efficient, as it generates random values for each event, even as the acceptance probability, $\frac{n}{i}$, converges to 0. 

Algorithm L also constructs a uniform sample over a stream, but improves on Algorithm R's performance by generating significantly fewer random values on large streams.

## Use

```go
import "github.com/dmw2151/reservoir"

func main() {
	// New Sampler w. reservoir of 4 events, rand seed of 2152
	s := reservoir.NewReservoirSample[int](4, 2152)

	// Feed Sampler 10 events
	for i := 0; i < 10; i++ {
		s.ReadSample(i)
	}

	// Check Sampler's reservoir, retains: [3 7 8 2]
	fmt.Printf("reservoir: %d\n" , s.Samples())
}
```

## Benchmarks

```text
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
Benchmark_EvaluateSampler_AlgorithmL/1M_tiny_reservoir-8    92446       12619 ns/op
Benchmark_EvaluateSampler_AlgorithmL/1M_small_reservoir-8   23392       50853 ns/op
Benchmark_EvaluateSampler_AlgorithmL/8M_tiny_reservoir-8       86    14206785 ns/op
Benchmark_EvaluateSampler_AlgorithmL/8M_small_reservoir-8      86    14653064 ns/op
Benchmark_EvaluateSampler_AlgorithmL/8M_large_reservoir-8      39    28500859 ns/op
```

```text
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
Benchmark_EvaluateSampler_AlgorithmR                                        
Benchmark_EvaluateSampler_AlgorithmR/1M_tiny_reservoir-8    50774       23666 ns/op
Benchmark_EvaluateSampler_AlgorithmR/1M_small_reservoir-8   58366       20586 ns/op
Benchmark_EvaluateSampler_AlgorithmR/8M_tiny_reservoir-8        6   174645414 ns/op
Benchmark_EvaluateSampler_AlgorithmR/8M_small_reservoir-8       6   177875939 ns/op
Benchmark_EvaluateSampler_AlgorithmR/8M_large_reservoir-8       6   180855144 ns/op
```

## Appendix

...

## Additional Reading

- *Random Sampling With a Reservoir* ([Vitter, 1985](https://www.cs.umd.edu/~samir/498/vitter.pdf))

- *Non-Uniform Random Variate Generation* ([Devroye, 1986](http://luc.devroye.org/chapter_twelve.pdf))