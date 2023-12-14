package pokeapi

import (
	"sync"
	"time"
)

type PokeCache struct {
	entries map[string]pokeCacheEntry
	mutex *sync.Mutex
}

type pokeCacheEntry struct{
	data []byte
	createdAt time.Time
}

func NewCache(interval, staleTime time.Duration) PokeCache {
	cache := PokeCache{
		entries: map[string]pokeCacheEntry{},
		mutex: &sync.Mutex{},
	}

	go cache.reapLoop(interval, staleTime)

	return cache
}

func (cache *PokeCache) reapLoop(interval, staleTime time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(staleTime)		
	}
}

func (cache *PokeCache) reap(staleTime time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	
	now := time.Now().UTC()
	for key, entry := range cache.entries {
		if entry.createdAt.Before(now.Add(-staleTime)) {
			delete(cache.entries, key)
		}
	}
}

func (cache *PokeCache) Get(url *string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	entry, ok := cache.entries[*url]
	if !ok {
		return []byte{}, false
	}

	return entry.data, true
}

func (cache *PokeCache) Add(url string, data []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.entries[url] = pokeCacheEntry{
		data: data,
		createdAt: time.Now().UTC(),
	}
}