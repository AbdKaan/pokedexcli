package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu         sync.Mutex
	cacheEntry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{cacheEntry: make(map[string]cacheEntry)}
	go cache.reapLoop(duration)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntry[key] = cacheEntry{time.Now(), val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	result, found := c.cacheEntry[key]
	if found {
		return result.val, found
	} else {
		return []byte{}, found
	}
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		c.mu.Lock()
		// Find and delete caches that are older than the given duration
		for key, cache := range c.cacheEntry {
			if t.Sub(cache.createdAt) > duration {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
