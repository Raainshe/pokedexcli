package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	Mu       sync.Mutex
	Interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	var newCache Cache

	newCache.Entries = make(map[string]cacheEntry)
	newCache.Interval = interval
	go newCache.reapLoop()
	return &newCache
}

func (c *Cache) Add(url string, data []byte) {
	//check if it exists first
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entries[url] = cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if value, exists := c.Entries[url]; !exists {
		return nil, false
	} else {
		return value.val, true
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.Mu.Lock()
			for url, entry := range c.Entries {
				if time.Since(entry.createdAt) > c.Interval {
					delete(c.Entries, url)
				}
			}
			c.Mu.Unlock()
		}
	}

}
