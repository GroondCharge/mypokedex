package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}

/*

Funkcija newcache bo ustvarla nov Cache struct, podamo trenutni čas
Map vrednosti bo
- key: url string
- value: slice of bytes k ga vrne request

Potem bom preverjal
resut_bytes, ok := Cache.val[url]
if !ok{
	potem rezultat ni v cachu, requestaj nove
}else{
	serviraj iz cacha}


type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	cahceEntries map[string]cacheEntry
	mu           sync.Mutex
}
*/
