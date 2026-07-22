package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu      sync.RWMutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:      sync.RWMutex{},
		entries: make(map[string]cacheEntry),
	}
	ticker := time.NewTicker(interval)
	go cache.reapLoop(ticker, interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = e
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e, ok := c.entries[key]
	if !ok {
		return []byte{}, ok
	}
	return e.val, true
}

func (c *Cache) reapLoop(ticker *time.Ticker, interval time.Duration) {
	for range ticker.C {
		c.mu.Lock()
		for k, e := range c.entries {
			if e.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
