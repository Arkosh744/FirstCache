package FirstCache

import (
	"errors"
	"sync"
)

type cache struct {
	sync.RWMutex
	data map[string]any
}

func NewCache() *cache {
	return &cache{data: make(map[string]any)}
}

func (cache *cache) Set(key string, value any) {
	cache.Lock()
	defer cache.Unlock()

	cache.data[key] = value
}

func (cache *cache) Get(key string) (any, error) {
	cache.RLock()
	defer cache.RUnlock()

	value, ok := cache.data[key]
	if !ok {
		return value, errors.New("key doesnt exist")
	}
	return value, nil
}

func (cache *cache) Delete(key string) error {
	cache.Lock()
	defer cache.Unlock()

	_, ok := cache.data[key]
	if ok {

		delete(cache.data, key)
		return nil
	}

	return errors.New("key doesnt exist")
}
