# v2.1.2 Benchmark Baseline

Date: 2026-05-20

Environment:

- Go: `go version go1.26.3-X:nodwarf5 linux/amd64`
- OS/arch: `Linux x86_64`
- Kernel: `7.0.7-arch1-1`
- CPU: `AMD Ryzen 7 3700X 8-Core Processor`

Previous release:

- Ref: `v2.1.1`
- Command: `GOWORK=off go test -run=^$ -bench=. -benchmem ./...`
- Result: command passed, but `v2.1.1` did not include benchmark functions or
  the `make bench` target, so there is no numeric baseline to compare against.

Release branch:

- Ref: `v2.1.2` release commit candidate
- Command: `make bench`
- Result: command passed.

Selected output:

```text
BenchmarkCryptoSourceRead-16     	  688524	      1741 ns/op	       0 B/op	       0 allocs/op
BenchmarkDerivedSourceRead-16    	  595362	      1891 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint64n-16              	31202192	        37.23 ns/op	       8 B/op	       1 allocs/op
BenchmarkFloat64-16              	36576406	        32.11 ns/op	       8 B/op	       1 allocs/op
BenchmarkBytes16-16              	21778358	        52.66 ns/op	      16 B/op	       1 allocs/op
BenchmarkNormal-16               	13264831	        91.57 ns/op	       8 B/op	       1 allocs/op
BenchmarkPoisson-16              	 1550178	       765.0 ns/op	     103 B/op	      12 allocs/op
BenchmarkGamma-16                	 7693809	       154.0 ns/op	      16 B/op	       2 allocs/op
BenchmarkID-16                   	 3326094	       354.8 ns/op	     152 B/op	       2 allocs/op
BenchmarkString32-16             	 2720612	       436.7 ns/op	     160 B/op	       2 allocs/op
BenchmarkTokenHex32-16           	 9047937	       133.7 ns/op	      48 B/op	       2 allocs/op
BenchmarkTokenURLSafe32-16       	 7113224	       167.2 ns/op	      56 B/op	       2 allocs/op
BenchmarkULID-16                 	 6853886	       175.9 ns/op	      48 B/op	       2 allocs/op
BenchmarkV4-16                   	 7519795	       158.6 ns/op	      64 B/op	       2 allocs/op
BenchmarkV7-16                   	 5836332	       205.5 ns/op	      64 B/op	       2 allocs/op
```

No performance improvement or regression claim is made because the previous
release does not contain comparable benchmark functions.
