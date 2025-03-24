package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mu *sync.Mutex
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
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	// TODO: Kick off readLoop to remove old entries
	cache := Cache{
		mu: &sync.Mutex{},
		cache: make(map[string]cacheEntry),
		
	}

	go cache.readLoop(interval)
	return cache
}