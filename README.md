# TTL cache
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
keep that in mind.

## Licence
Copyright (c) 2016 Johan Sageryd <j@1616.se>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
