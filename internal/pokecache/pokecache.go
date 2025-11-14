package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}
type Cache struct{
	mutex sync.RWMutex
	entries map[string]cacheEntry
	interval time.Duration
}



func NewCache(reapInterval time.Duration) *Cache{
	c := &Cache{
		entries : make(map[string]cacheEntry),
		interval : reapInterval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add (key string, val []byte){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:val,
	}
}


func (c *Cache) Get(key string)([]byte, bool){
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, ok := c.entries[key]
	if !ok{
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(){
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C{
		c.mutex.Lock()
		for key, entry := range c.entries{
			if time.Since(entry.createdAt) > c.interval{
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}

}
