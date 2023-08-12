package main

import (
	"sync"
	"time"
)

/*
This is an experiment to create an in-memory store like redis using Go
*/

func main() {
	// rd := NewRedisStore()

	// ttl := 1 * time.Second
	// ttl2 := 7 * time.Second
	// rd.Set("test key", "Test value", ttl)
	// rd.Set("test key 2", "Test value 2", ttl2)
	// fmt.Printf("%+v\n", rd.items)
	// time.Sleep(5 * time.Second)
	// fmt.Printf("%+v\n", rd.items)

	// PQ version
	PriorityQCache()
}

// Create the struct that represents the in-memory store
type RedisStore struct {
	items map[string]value
	mutex sync.Mutex
}

func NewRedisStore() *RedisStore {
	return &RedisStore{
		items: make(map[string]value),
		mutex: sync.Mutex{},
	}
}

type value struct {
	value string
	ttl   time.Duration
}

// Create methods to expose Get, Set, Update and Delete features
func (r *RedisStore) Set(k, v string, t time.Duration) {
	value := value{
		value: v,
		ttl:   t,
	}

	r.items[k] = value

	// In a separate thread, start a Goroutine to delete the item from the map when the ttl expires
	go r.TTLEviction(k, t)
}

// Create a TTL method to delete a specific key:value when the TTL expires
func (r *RedisStore) TTLEviction(k string, ttl time.Duration) {
	<-time.After(ttl)
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.items, k)
}
