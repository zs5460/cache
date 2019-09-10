package cache

import (
	"sync"
	"time"
)

type item struct {
	value   interface{}
	expires int64
}

// Cache stores data with expiration time.
type Cache struct {
	TTL time.Duration
	sync.Map
	close chan struct{}
}

// Close closes the cache.
func (c *Cache) Close() {
	c.close <- struct{}{}
	c.Map = sync.Map{}
}

// Get gets the value for the key.
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	v, ok := c.Map.Load(key)
	if !ok {
		return nil, false
	}

	if v.(item).expires < time.Now().UnixNano() {
		return nil, false
	}
	return v.(item).value, true
}

// Set sets a value for the key with an expiration time.
func (c *Cache) Set(key, value interface{}) {
	c.Map.Store(key, item{
		value:   value,
		expires: time.Now().Add(c.TTL).UnixNano(),
	})
}

// New returns a cache.
func New(ttl time.Duration) *Cache {
	c := &Cache{
		close: make(chan struct{}),
		TTL:   ttl,
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-c.close:
				return

			case <-ticker.C:
				now := time.Now().UnixNano()
				c.Map.Range(func(k, v interface{}) bool {
					if v.(item).expires > now {
						c.Map.Delete(k)
					}
					return true
				})
			}
		}
	}()

	return c
}
