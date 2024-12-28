# TTL cache

[![Build Status](https://github.com/jsageryd/ttlcache/workflows/ci/badge.svg)](https://github.com/jsageryd/ttlcache/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsageryd/ttlcache)](https://goreportcard.com/report/github.com/jsageryd/ttlcache)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jsageryd/ttlcache)
[![Licence MIT](https://img.shields.io/badge/licence-MIT-lightgrey.svg?style=flat)](https://github.com/jsageryd/ttlcache#licence)

Simple thread-safe expiring in-memory cache.

## Usage example
```go
package main

import (
	"fmt"
	"time"

	"github.com/jsageryd/ttlcache"
)

func main() {
	cache := ttlcache.New(5 * time.Second)
	cache.Set("foo", "bar")
	if value, ok := cache.Get("foo"); ok {
		fmt.Println("Value:", value) // (type assert value as needed)
	} else {
		fmt.Println("Key not found")
	}
}
```

## Note
For every new key that is set a goroutine is spawned to expire that key after
the timeout. This keeps the implementation clean and simple. Already existing
keys keep the same goroutine if they are set again before they have expired.
While it should not generally be an issue to spawn many thousand goroutines, if
you have a very large amount of things to cache, this package is not for you. So
keep that in mind. Consider [Qcache](https://github.com/jsageryd/qcache).
