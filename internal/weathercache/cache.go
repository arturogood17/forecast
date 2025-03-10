package weathercache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]ValSet
	mu    *sync.Mutex
}

type ValSet struct {
	createdAt time.Time
	value     []byte
}

func WCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]ValSet),
		mu:    &sync.Mutex{},
	}
	go c.reapLoopC(interval)

	return c
}

func (c *Cache) AddC(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = ValSet{
		createdAt: time.Now().UTC(),
		value:     val,
	}
}

func (c *Cache) GetC(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, exist := c.cache[key]
	return v.value, exist
}

func (c *Cache) reapLoopC(interval time.Duration) {
	t := time.NewTicker(interval)
	for range t.C {
		c.reapC(time.Now().UTC(), interval)
	}
}

func (c *Cache) reapC(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, k)
		}
	}
}
