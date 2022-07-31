package FirstCache

import (
	"errors"
	"sync"
)

type cache struct {
	mu   *sync.RWMutex
	data map[string]any
}

func NewCache() *cache {
	return &cache{data: make(map[string]any)}
}

func (cache *cache) Set(key string, value any) {
	cache.mu.Lock()
	cache.data[key] = value
	cache.mu.Unlock()

}

func (cache *cache) Get(key string) (any, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	value, ok := cache.data[key]
	if !ok {
		return value, errors.New("key doesnt exist")
	}
	return value, nil
}

func (cache *cache) Delete(key string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	_, ok := cache.data[key]
	if ok {

		delete(cache.data, key)
		return nil
	}

	return errors.New("key doesnt exist")
}
