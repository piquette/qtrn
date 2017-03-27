[![GoDoc](https://godoc.org/github.com/lucy/runewidth?status.svg)](https://godoc.org/github.com/lucy/runewidth)

This package is primarily based on https://github.com/jquast/wcwidth and
codegen from https://github.com/golang/text.

Benchmark improvements vs [github.com/mattn/go-runewidth][1]:
```
benchmark                 old ns/op     new ns/op     delta
BenchmarkEasyRune-8       156           13.3          -91.47%
BenchmarkEasyString-8     2614          83.0          -96.82%
Benchmark1-8              159           20.1          -87.36%
Benchmark2-8              1062          66.8          -93.71%
Benchmark3-8              7180          432           -93.98%
Benchmark4-8              147057        8834          -93.99%
```

[1]: https://github.com/mattn/go-runewidth
