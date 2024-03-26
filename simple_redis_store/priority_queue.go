package main

import (
	"container/heap"
	"fmt"
	"log"
	"time"
)

type Redis struct {
	items      map[string]*values
	expiration *ExpirationHeap
}

type values struct {
	value   interface{}
	ttl     time.Duration
	expires time.Time
}

func NewCache() *Redis {
	store := make(map[string]*values)
	expiration := &ExpirationHeap{}
	heap.Init(expiration)
	return &Redis{
		items:      store,
		expiration: expiration,
	}
}

// Gets a key:value from the cache. It checks if the value's expired time is after the current time before returning the key
func (r *Redis) get(k string) *values {
	now := time.Now()
	if val, ok := r.items[k]; ok {
		if now.After(val.expires) {
			log.Println("🚨 KEY HAS EXPITED IN GET")
			return nil
		}
	} else {
		fmt.Println("🚨 KEY IS NOT IN MEMORY: ", k)
		return nil
	}
	return r.items[k]
}

func (r *Redis) set(k string, v interface{}, t time.Duration) {
	ttl := t * time.Second
	expTime := time.Now().Add(ttl)

	newValues := values{
		value:   v,
		ttl:     t,
		expires: expTime,
	}

	// Set the k:v pair into the Redis struct
	r.items[k] = &newValues

	// Push the key of the new item into the min heap data structure. This is used later in the interval to delete those expired items from the Redis struct items map
	heap.Push(r.expiration, &expirationItem{
		key:      k,
		expireAt: expTime,
	})
}

// Loops through the min heap, once it finds a item whose exp time is later than the current time, it stops the loop
func (r *Redis) removeExpired() {
	now := time.Now()
	// As long as there are items in the min heap, we keep going
	for r.expiration.Len() > 0 {
		// At each iteration, we pop the min value from the heap ({key, expiration Time})
		item := heap.Pop(r.expiration).(*expirationItem)

		// If current time is < the min heap's min value's expired time, we push the item back into the heao and break the loop. This signifies that the min item in the heap is still not expired (meaning that the rest of the items are also still valid and should be kept)
		if now.Before(item.expireAt) {
			heap.Push(r.expiration, item)
			break
		}
		// Perform deletion of k:v from the Redis struct storing all k:v pairs
		fmt.Println("DELETING KEY: ", item.key)
		delete(r.items, item.key)
	}
}

// Every 2 seconds, the clean cycle will begin. Removing expired keys from the cache
func (r *Redis) StartCleanInterval() {

	ticker := time.NewTicker(2 * time.Second)
	// quit := make(chan struct{})
	go func() {
		defer ticker.Stop()
		for t := range ticker.C {
			fmt.Println("⏰ t running: ", t)
			r.removeExpired()
		}
		// for {
		// 	select {
		// 	case <-ticker.C:
		// 		// do stuff
		// 		r.removeExpired()
		// 	}
		// }
	}()
}

func PriorityQCache() {

	// Initialising a new cache instance
	cache := NewCache()
	cache.StartCleanInterval()

	// Setting values in cache
	cache.set("test key", "test value", time.Duration(2))
	cache.set("test key1", "test value", time.Duration(2))
	cache.set("test key2", "test value", time.Duration(3))
	cache.set("test key3", "test value", time.Duration(4))

	// Getting a value for testing
	res := cache.get("test key")
	fmt.Println("result 'test key':", res.value.(string))
	fmt.Println("results:", cache.items)

	// Simulate running time
	time.Sleep(5 * time.Second)

	// Get key from cache to test
	exp := cache.get("test key")
	fmt.Println("EXP KEY: ", exp)
	// cache.removeExpired()

	r := cache.get("test key")
	fmt.Println("🚨:", r)
	fmt.Println("results:", cache.items)
}

type expirationItem struct {
	key      string
	expireAt time.Time
}

type ExpirationHeap []*expirationItem

func (h ExpirationHeap) Len() int           { return len(h) }
func (h ExpirationHeap) Less(i, j int) bool { return h[i].expireAt.Before(h[j].expireAt) }
func (h ExpirationHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ExpirationHeap) Push(x interface{}) {
	*h = append(*h, x.(*expirationItem))
}

func (h *ExpirationHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func (h ExpirationHeap) Peek() *expirationItem {
	if len(h) == 0 {
		return nil
	}
	return h[0]
}
