# Swish

An optimized Swish activation function ([Ramachandran, Zoph and Le, 2017](https://arxiv.org/abs/1710.05941)), for neural networks.

## Screenshots

Output from the graph example in `cmd/graph`:

![](img/swish.png)

![](img/sigmoid.png)

The graphs are drawn using [goterm](https://github.com/buger/goterm).

## Benchmark Results

### Using a `Swish` function that uses `math.Exp`

First run:

```
goos: linux
goarch: amd64
pkg: github.com/xyproto/swish
BenchmarkSwish07-8   	200000000	         8.93 ns/op
BenchmarkSwish03-8   	200000000	         8.95 ns/op
PASS
ok  	github.com/xyproto/swish	5.391s
```

### Using the optimized `Swish` function that uses `exp256`

```
goos: linux
goarch: amd64
pkg: github.com/xyproto/swish
BenchmarkSwish07-8   	2000000000	         0.26 ns/op
BenchmarkSwish03-8   	2000000000	         0.26 ns/op
PASS
ok  	github.com/xyproto/swish	1.108s
```

The optimized `Swish` function is *34x* faster than the one that uses `math.Exp`.

The average difference between the optimized and non-optimzed version is `+-0.0013` and the maximum difference in precision is `+-0.0024`. This is when `x` is in the range `[5,3]`. See the program in `cmd/precision` for how this was calculated.

## General info

* Version: 2.0.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
