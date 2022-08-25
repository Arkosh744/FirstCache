package FirstCache

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

type Cache struct {
	data map[string]Value
	mu   sync.RWMutex
	ttl  time.Duration
}

type Value struct {
	Value any
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]Value)}
}

func (cache *Cache) Set(key string, value any, ctx context.Context) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.data[key] = Value{value}

	go cache.killElement(ctx, key)
}

func (cache *Cache) Get(key string) (Value, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	value, ok := cache.data[key]
	if !ok {
		return value, errors.New("Cannot Get this key - key doesnt exist")
	}
	return value, nil
}

func (cache *Cache) Delete(key string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	_, ok := cache.data[key]
	if ok {
		delete(cache.data, key)
		return nil
	}
	return errors.New("Cannot delete key - key doesnt exist")
}

func (cache *Cache) killElement(ctx context.Context, key string) {
	value, _ := cache.Get(key)

	ctx, _ = context.WithTimeout(ctx, cache.ttl)

	select {
	case <-ctx.Done():
		return

	case <-time.After(cache.ttl):
		err := cache.Delete(key)
		if err == nil {
			log.Println("ttl timeout AFTER for value ", value)
			return
		}
		log.Println(err)
		return
	}

}
