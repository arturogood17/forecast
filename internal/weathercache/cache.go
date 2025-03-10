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
	go c.ReapLoopC(interval)

	return c
}

func (c *Cache) addC(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = ValSet{
		createdAt: time.Now().UTC(),
		value:     val,
	}
}

func (c *Cache) getC(key string) any {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cache[key].value
}

func (c *Cache) ReapLoopC(interval time.Duration) {
	t := time.NewTicker(interval)
	for t.C {
		c.ReapC(time.Now().UTC(), interval)
	}
}

func (c *Cache) ReapC(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, k)
		}
	}
}
