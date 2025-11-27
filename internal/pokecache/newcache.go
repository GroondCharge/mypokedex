package pokecache

import "time"

func NewCache(interval time.Time) {

}
func (c *Cache) Add()      {}
func (c *Cache) Get()      {}
func (c *Cache) reapLoop() {}

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
