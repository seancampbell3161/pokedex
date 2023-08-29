package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		make(map[string]cacheEntry),
		&sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) AddToCache(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) GetCache(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheItem, ok := c.entries[key]
	if len(cacheItem.value) == 0 {
		ok = false
	}

	return cacheItem.value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	timer := time.NewTicker(interval)
	for range timer.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(current time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		if current.Sub(entry.createdAt) > last {
			delete(c.entries, key)
		}
	}
}
