# RLP vs JSON Encoding/Decoding Performance Comparison

This repository contains the code used to compare the performance of RLP and JSON encoding/decoding in Golang.

Run it with:

```bash
go test -bench=. -benchmem
```

Results as below:

```bash
goos: linux
goarch: amd64
pkg: rlp-vs-json
cpu: AMD Ryzen 7 5700G with Radeon Graphics
BenchmarkJSONEncoding-16                    4483            258600 ns/op          189691 B/op        990 allocs/op
BenchmarkRLPEncoding-16                    33832             35169 ns/op           49815 B/op          2 allocs/op
BenchmarkJSONDecoding-16                    1846            664327 ns/op          122480 B/op       1101 allocs/op
BenchmarkRLPDecoding-16                     7880            146389 ns/op          150855 B/op       1349 allocs/op
BenchmarkStorageComparison-16           1000000000               0.001566 ns/op        0 B/op          0 allocs/op
--- BENCH: BenchmarkStorageComparison-16
    call_test.go:136: JSON size: 92180 bytes
    call_test.go:137: RLP size: 41640 bytes
        ... [output truncated]
PASS
ok      rlp-vs-json     5.230s
```
