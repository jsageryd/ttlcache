// Package ttlcache provides a cache whose items expire after set TTL.
package ttlcache

import (
	"sync"
	"time"
)

// Cache is a key-value cache for arbitrary values
type Cache interface {
	Expire(key interface{})
	ExpireAll()
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{})
}

type item struct {
	expired chan struct{}
	timer   *time.Timer
	value   interface{}
}

type cache struct {
	sync.RWMutex

	items map[interface{}]item
	ttl   time.Duration
}

// New returns a new cache. ttl is the time to live for its items.
func New(ttl time.Duration) Cache {
	c := cache{
		items: make(map[interface{}]item),
		ttl:   ttl,
	}
	return &c
}

// Expire deletes given key.
func (c *cache) Expire(key interface{}) {
	c.Lock()
	if i, ok := c.items[key]; ok {
		close(i.expired)
	}
	delete(c.items, key)
	c.Unlock()
}

// ExpireAll expires all keys.
func (c *cache) ExpireAll() {
	c.Lock()
	for _, i := range c.items {
		close(i.expired)
	}
	c.items = make(map[interface{}]item)
	c.Unlock()
}

// Get retrieves a value for given key. Returned boolean is true if value
// exists, otherwise false.
func (c *cache) Get(key interface{}) (interface{}, bool) {
	c.RLock()
	if i, ok := c.items[key]; ok {
		c.RUnlock()
		return i.value, true
	}
	c.RUnlock()
	return nil, false
}

// Set updates the cache to store given value at given key.
func (c *cache) Set(key interface{}, value interface{}) {
	c.Lock()
	defer c.Unlock()
	if i, ok := c.items[key]; ok {
		i.timer.Reset(c.ttl)
		i.value = value
	} else {
		i = item{
			expired: make(chan struct{}),
			timer:   time.NewTimer(c.ttl),
			value:   value,
		}
		c.items[key] = i
		go func() {
			select {
			case <-i.timer.C:
				c.Expire(key)
			case <-i.expired:
				return
			}
		}()
	}
}
