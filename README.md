# cache

[![Build Status](https://travis-ci.org/zs5460/cache.svg?branch=master)](https://travis-ci.org/zs5460/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/zs5460/cache)](https://goreportcard.com/report/github.com/zs5460/cache)
[![codecov](https://codecov.io/gh/zs5460/cache/branch/master/graph/badge.svg)](https://codecov.io/gh/zs5460/cache)
[![GoDoc](https://godoc.org/github.com/zs5460/cache?status.svg)](https://godoc.org/github.com/zs5460/cache)

## Install

```shell
go get -u github.com/zs5460/cache
```

## Usage

```go
c := cache.New(1*time.Minute)

c.Set("foo","bar")

v , exist := c.Get("foo")
if exist {
    fmt.Println(v.(string))
}

c.Close()
```

## Benchmark

```shell
goos: windows
goarch: amd64
pkg: github.com/zs5460/cache
BenchmarkNew-4      	  3000000	         560 ns/op	     368 B/op	       5 allocs/op
BenchmarkGet-4      	 30000000	        36.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet-4      	  3000000	         414 ns/op	      48 B/op	       2 allocs/op
BenchmarkDelete-4   	100000000	        11.6 ns/op	       0 B/op	       0 allocs/op
```

## License

Released under MIT license, see [LICENSE](LICENSE) for details.
