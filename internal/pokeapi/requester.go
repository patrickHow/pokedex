package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type RequestManager struct {
	cache *pokecache.Cache
}

const defaultCacheDuration time.Duration = 30

func NewRequestManager() *RequestManager {
	return &RequestManager{
		cache: pokecache.NewCache(defaultCacheDuration),
	}
}

func (rqm *RequestManager) GetData(url string) ([]byte, error) {

	// First see if the data is in the cache
	t0 := time.Now()
	data, ok := rqm.cache.Get(url)

	if ok {
		t1 := time.Now()
		fmt.Printf("Cache hit on key %s\n", url)
		fmt.Printf("Rq duration: %v\n", t1.Sub(t0))
		return data, nil
	}

	// If data was not in the cache, make an HTTP request
	fmt.Printf("Cache miss on key %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-ok HTTP status: %s", resp.Status)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// Cache data before returning
	rqm.cache.Add(data, url)
	t1 := time.Now()
	fmt.Printf("Rq duration: %v\n", t1.Sub(t0))

	return data, nil
}
