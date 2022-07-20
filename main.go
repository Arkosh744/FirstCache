package FirstCache

import "errors"

type cache struct {
	data map[string]interface{}
}

func NewCache() *cache {
	return &cache{data: make(map[string]interface{})}
}

func (cache *cache) Set(key string, value interface{}) {
	cache.data[key] = value
}

func (cache *cache) Get(key string) (interface{}, error) {
	value, ok := cache.data[key]
	if !ok {
		return value, errors.New("key doesnt exist")
	}
	return value, nil
}

func (cache *cache) Delete(key string) {
	delete(cache.data, key)
}
