package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type cache struct {
	data map[string]Value
	mu   sync.RWMutex
	ctx  context.Context
}

type Value struct {
	Value any
	ttl   time.Duration
}

func NewCache() *cache {
	return &cache{data: make(map[string]Value), ctx: context.Background()}
}

func (cache *cache) Set(key string, value any, ttl time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.data[key] = Value{value, ttl}

	go cache.killElement(cache.ctx, key)
}

func (cache *cache) Get(key string) (Value, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	value, ok := cache.data[key]
	if !ok {
		return value, errors.New("Cannot Get this key - key doesnt exist")
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
	return errors.New("Cannot delete key - key doesnt exist")
}

func (cache *cache) killElement(ctx context.Context, key string) {
	value, _ := cache.Get(key)

	ctx, _ = context.WithTimeout(ctx, value.ttl)

	select {
	case <-ctx.Done():
		return

	case <-time.After(value.ttl):
		err := cache.Delete(key)
		if err == nil {
			fmt.Println("ttl timeout AFTER for ", value)
			return
		}
		fmt.Println(err)
		return
	}

}

func main() {
	cache := NewCache()

	cache.Set("userId", 42, time.Second*5)
	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	fmt.Println(userId.Value) // Output: 42

	//err = cache.Delete("userId")
	if err != nil {
		return
	}
	time.Sleep(time.Second * 6) // прошло 5 секунд
	userId, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}

}
