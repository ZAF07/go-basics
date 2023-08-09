package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Redis struct {
	items      map[string]values
	expiration *ExpirationHeap
}

func NewCache() *Redis {
	store := make(map[string]values)
	expiration := &ExpirationHeap{}
	heap.Init(expiration)
	return &Redis{
		items:      store,
		expiration: expiration,
	}
}

type values struct {
	value interface{}
	ttl   time.Duration
}

func (r *Redis) get(k string) values {
	return r.items[k]
}

func (r *Redis) set(k string, v interface{}, t time.Duration) {
	ttl := t * time.Second

	newValues := values{
		value: v,
		ttl:   t,
	}

	r.items[k] = newValues
	expTime := time.Now().Add(ttl)
	heap.Push(r.expiration, &expirationItem{
		key:      k,
		expireAt: expTime,
	})
}

func (r *Redis) removeExpired() {
	now := time.Now()
	for r.expiration.Len() > 0 {
		item := heap.Pop(r.expiration).(*expirationItem)
		if now.Before(item.expireAt) {
			heap.Push(r.expiration, item)
			break
		}
		fmt.Println("Expired key:", item.key)
		delete(r.items, item.key)
	}
}

func Testing() {
	fmt.Println("testing")
}

func PriorityQCache() {
	cache := NewCache()
	cache.set("test key", "test value", time.Duration(2))
	res := cache.get("test key")
	fmt.Println("result:", res)
	time.Sleep(3 * time.Second)
	cache.removeExpired()
	r := cache.get("test key")
	fmt.Println("🚨:", r)
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
