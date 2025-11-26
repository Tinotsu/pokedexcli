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

type cache struct {
	entry map[string]cacheEntry
	mu sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val []byte
}
func NewCache(interval time.Duration) *cache {
	c := new(cache)
	reapLoop(c, interval)
	return c
}
func (c *cache) Add(key string, val []byte) {
	c.mu.Lock()
	m := new(cacheEntry)
	m.val = val
	c.entry[key] = *m
	c.mu.Unlock()
} 
func (c *cache) Get(key string) ([]byte, bool) {
	val := c.entry[key].val
	if val != nil {
		return val, true
	}
	return []byte{}, false
}
func reapLoop(c *cache, interval time.Duration) {
	c.mu.Lock()
	for k, v := range c.entry {
		t1 := time.Now()
		if t1.Sub(v.createdAt) > interval {
			delete(c.entry, k)
		}
	}
	c.mu.Unlock()
}
