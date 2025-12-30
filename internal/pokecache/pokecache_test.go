package pokecache

import (
	"testing"
	"time"
)

// TestAddGet verifies that adding and retrieving a cache entry works correctly.
func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "https://pokeapi.co/api/v2/location-area"
	val := []byte("test data")

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key %s", key)
		return
	}
	if string(got) != string(val) {
		t.Errorf("expected %s, got %s", string(val), string(got))
	}
}

// TestGetMissing verifies that getting a non-existent key returns false.
func TestGetMissing(t *testing.T) {
	cache := NewCache(5 * time.Second)

	_, ok := cache.Get("nonexistent")
	if ok {
		t.Error("expected to not find nonexistent key")
	}
}

// TestMultiplePages verifies that multiple pages can be cached and retrieved (simulating map/mapb navigation).
func TestMultiplePages(t *testing.T) {
	cache := NewCache(5 * time.Second)

	pages := []struct {
		url  string
		data string
	}{
		{"https://pokeapi.co/api/v2/location-area?offset=0&limit=20", "page 1 data"},
		{"https://pokeapi.co/api/v2/location-area?offset=20&limit=20", "page 2 data"},
		{"https://pokeapi.co/api/v2/location-area?offset=40&limit=20", "page 3 data"},
	}

	// Add all pages to cache (simulating forward navigation with map)
	for _, page := range pages {
		cache.Add(page.url, []byte(page.data))
	}

	// Verify all pages are cached (simulating backward navigation with mapb)
	for _, page := range pages {
		got, ok := cache.Get(page.url)
		if !ok {
			t.Errorf("expected to find cached page %s", page.url)
			continue
		}
		if string(got) != page.data {
			t.Errorf("expected %s, got %s", page.data, string(got))
		}
	}
}

// TestReapLoop verifies that old cache entries are automatically removed after the interval.
func TestReapLoop(t *testing.T) {
	interval := 50 * time.Millisecond
	cache := NewCache(interval)

	key := "https://pokeapi.co/api/v2/location-area"
	val := []byte("test data")

	cache.Add(key, val)

	// Should exist immediately
	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key %s", key)
		return
	}

	// Wait for reap to happen
	time.Sleep(interval * 3)

	// Should be gone after reap
	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected key %s to be reaped", key)
	}
}
