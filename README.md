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
goos: linux
goarch: amd64
pkg: github.com/zs5460/cache
BenchmarkNew-8      	  10000000	          188 ns/op	     368 B/op	       5 allocs/op
BenchmarkGet-8      	 100000000	         18.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet-8      	  10000000	          209 ns/op	      48 B/op	       2 allocs/op
BenchmarkDelete-8   	1000000000	         3.13 ns/op	       0 B/op	       0 allocs/op
```

## License

Released under MIT license, see [LICENSE](LICENSE) for details.
