// Package pokecache cache helper
package pokecache

import(
	"fmt"
	"time"
	"sync"
)

func CacheTest() {
	fmt.Print("pokecache")
}
type Cache struct {
	entry map[string]cacheEntry
	Mu sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val []byte
}
func reapLoop(c *Cache, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		timeNow := time.Now()
		c.Mu.Lock()
		for k, v := range c.entry {
			if timeNow.Sub(v.createdAt) > interval {
				delete(c.entry, k)
			}
		}
		c.Mu.Unlock()
	}
}
func NewCache(interval time.Duration) *Cache {
	c := new(Cache)
	c.entry = map[string]cacheEntry{}
	go reapLoop(c, interval)
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	existsKey := false
	for k := range c.entry {
		if k == key {
			existsKey = true
		}
	}
	if !existsKey {
		m := new(cacheEntry)
		m.val = val
		m.createdAt = time.Now()
		c.entry[key] = *m
	} else {
		v := c.entry[key]
		v.createdAt = time.Now()
		v.val = val
		c.entry[key] = v
	}
	c.Mu.Unlock()
} 
func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	val := c.entry[key].val
	if val != nil {
		c.Mu.Unlock()
		return val, true
	}
	c.Mu.Unlock()
	return []byte{}, false
}
