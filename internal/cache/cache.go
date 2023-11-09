package cache

import (
	"errors"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	lock    sync.RWMutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		lock:    sync.RWMutex{},
	}
	cache.reapLoop(interval)
	return &cache
}

func (cache *Cache) Add(key string, value []byte) error {
	if key == "" {
		return errors.New("Key was not provided to Cache.Add method")
	}
	if len(value) == 0 {
		return errors.New("There's nothing to cache - value is empty")
	}

	cache.lock.Lock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
	cache.lock.Unlock()

	return nil
}

func (cache *Cache) Get(key string) (value []byte, ok bool) {
	cache.lock.RLock()
	requestedValue, ok := cache.entries[key]
	cache.lock.RUnlock()

	return requestedValue.value, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			currentTime := <-ticker.C
			for key, value := range cache.entries {
				if currentTime.Sub(value.createdAt) > interval {
					cache.lock.Lock()
					delete(cache.entries, key)
					cache.lock.Unlock()
				}
			}
		}
	}()
}
