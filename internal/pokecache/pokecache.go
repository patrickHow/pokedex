package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	raw       []byte
}

type Cache struct {
	data           map[string]cacheEntry
	deleteInterval time.Duration
	lock           *sync.RWMutex
}

func NewCache(deleteInterval time.Duration) *Cache {

	cache := Cache{
		data:           make(map[string]cacheEntry),
		deleteInterval: deleteInterval,
		lock:           &sync.RWMutex{},
	}

	c := time.Tick(deleteInterval * time.Second)
	go cache.reapLoop(c)

	return &cache
}

func (cache *Cache) Add(data []byte, key string) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	entry := cacheEntry{
		raw:       data,
		createdAt: time.Now(),
	}

	cache.data[key] = entry
	fmt.Printf("Caching data with key: %s\n", key)
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.lock.RLock()
	defer cache.lock.RUnlock()

	if data, ok := cache.data[key]; ok {
		return data.raw, true
	} else {
		return []byte{}, false
	}
}

func (cache *Cache) reapLoop(c <-chan time.Time) {
	for tm := range c {
		for key, data := range cache.data {
			if tm.Sub(data.createdAt) >= cache.deleteInterval {
				cache.lock.Lock()
				fmt.Printf("Deleting cache entry with key %s\n", key)
				delete(cache.data, key)
				cache.lock.Unlock()
			}
		}
	}
}
