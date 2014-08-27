benchmarker [![Build Status](https://drone.io/github.com/maciejmrowiec/benchmarker/status.png)](https://drone.io/github.com/maciejmrowiec/benchmarker/latest)
===========

Benchmarker is a simple utility to quicky compare results from golang benchmark suite between executions.

Usage:

old_results.log
```
Benchmark_FirstPage_LargeResult	           1	11981463000 ns/op
Benchmark_FirstPage_SmallResult	          10	 210276300 ns/op
Benchmark_SecondPage_SmallResult	      10	 210213400 ns/op
Benchmark_FirstPage_SmallResult_Sorted	  10	 218926600 ns/op
Benchmark_SecondPage_SmallResult_Sorted	  10	 204990600 ns/op
```

new_results.log
```
Benchmark_FirstPage_LargeResult	          10	 175612500 ns/op
Benchmark_FirstPage_SmallResult	          50	  30923420 ns/op
Benchmark_SecondPage_SmallResult	      50	  49082380 ns/op
Benchmark_FirstPage_SmallResult_Sorted	  20	  52893350 ns/op
Benchmark_SecondPage_SmallResult_Sorted	  50	  48749140 ns/op
```

benchmarker old_results.log new_results.log
```
Improvement	Benchmark
6822.67 %	benchmark_firstpage_largeresult
679.99 %	benchmark_firstpage_smallresult
413.90 %	benchmark_firstpage_smallresult_sorted
428.29 %	benchmark_secondpage_smallresult
420.50 %	benchmark_secondpage_smallresult_sorted
		
1753.07 %	Total improvement
```

Note: Results are sorted by benchmark name.