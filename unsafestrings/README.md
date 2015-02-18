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

