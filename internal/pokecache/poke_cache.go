package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mu sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cache_entry := cacheEntry{
		createdAt : time.Now(),
		val : value,
	}
	
	c.cache[key] = cache_entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	value, exists := c.cache[key]
	if exists {
		return value.val, true
	}
	return []byte{}, false
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			c.mu.Lock()

			now := time.Now()

			for k, v := range c.cache {
				if now.Sub(v.createdAt) > interval {
					delete(c.cache, k)
				}
			}

			c.mu.Unlock()
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	// TODO: Kick off readLoop to remove old entries
	cache := &Cache{
		cache: make(map[string]cacheEntry),
	}

	cache.readLoop(interval)
	return cache
}