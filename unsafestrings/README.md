# Unsafe Strings

This package is mostly here to ilustrate an [article in Norwgian](http://bepsays.com/2015/02/18/usikre-go-peikarar/), but the benchmarks can be interesting enough on their own.

`go test -test.run=NONE -bench=".*" -test.benchmem=true .`

On my laptop, this currently gives:

```
BenchmarkSafeBytesToString	 5000000	       246 ns/op	      48 B/op	       1 allocs/op
BenchmarkUnsafeBytesToString	2000000000	         1.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnsafeStringsReplacer	 3000000	       552 ns/op	       0 B/op	       0 allocs/op
BenchmarkSafeStringsReplacer	 2000000	       777 ns/op	      48 B/op	       1 allocs/op
BenchmarkMultipleBytesReplace	  500000	      2563 ns/op	     200 B/op	       9 allocs/op
BenchmarkMultiplesStringsReplace	  500000	      2807 ns/op	     288 B/op	       6 allocs/op
BenchmarkAppendString	100000000	        10.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendByteString	20000000	       114 ns/op	       8 B/op	       1 allocs/op
```

## Rerun 2016-10-12

Same benchmarks as above on Go 1.7.1 and a slightly faster PC:

```
▶ go test -test.run=NONE -bench=".*" -test.benchmem=true ./unsafestrings
BenchmarkSafeBytesToString-4         	30000000	        47.7 ns/op	      48 B/op	       1 allocs/op
BenchmarkUnsafeBytesToString-4       	2000000000	         1.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnsafeStringsReplacer-4     	 5000000	       247 ns/op	       0 B/op	       0 allocs/op
BenchmarkSafeStringsReplacer-4       	 5000000	       294 ns/op	      48 B/op	       1 allocs/op
BenchmarkMultipleBytesReplace-4      	 3000000	       412 ns/op	     144 B/op	       3 allocs/op
BenchmarkMultiplesStringsReplace-4   	 2000000	       648 ns/op	     288 B/op	       6 allocs/op
BenchmarkAppendString-4              	500000000	         3.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendByteString-4          	100000000	        13.0 ns/op	       0 B/op	       0 allocs/op
```
