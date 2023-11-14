package pokecache

//pokecache creates the cache for the pokeapi
import (
	"sync"
	"time"
)

// CacheEntry is the struct of the infromation bing stored in the Cache
type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

// Cache is the struct for pokeapi's cache
type Cache struct {
	cache map[string]CacheEntry
	mux   *sync.Mutex
}

// creates a new cache for pokeapi
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mux:   &sync.Mutex{},
	}
	c.reapLoop(interval)

	return Cache{}
}

// Add adds information the the cache
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

// Get is used to get information from the pokeapi
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.Val, ok
}

// reapLoop
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

// reap
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.CreatedAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
