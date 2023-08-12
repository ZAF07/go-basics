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

func NewCache() *Redis {
	store := make(map[string]*values)
	expiration := &ExpirationHeap{}
	heap.Init(expiration)
	return &Redis{
		items:      store,
		expiration: expiration,
	}
}

type values struct {
	value   interface{}
	ttl     time.Duration
	expires time.Time
}

func (r *Redis) get(k string) *values {
	now := time.Now()
	if val, ok := r.items[k]; ok {
		if now.After(val.expires) {
			log.Println("KEY HAS EXPITED")
			return nil
		}
	}
	// cacheItem := r.items[k]
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

	r.items[k] = &newValues
	heap.Push(r.expiration, &expirationItem{
		key:      k,
		expireAt: expTime,
	})
}

func (r *Redis) removeExpired() {
	now := time.Now()
	fmt.Println("LENGTH: ", r.expiration.Len())
	for r.expiration.Len() > 0 {
		fmt.Println("RUNNING REMOVEEXPIRE")
		item := heap.Pop(r.expiration).(*expirationItem)
		if now.Before(item.expireAt) {
			heap.Push(r.expiration, item)
			fmt.Println("BREALING")
			break
		}
		fmt.Println("Expired key:", item.key)
		delete(r.items, item.key)
	}
}

func testBreak() {
	nums := []int{1, 2, 3, 4}
	for len(nums) > 0 {
		i := nums[len(nums)-1]
		nums = append(nums[0 : len(nums)-1])
		fmt.Println("popped: ", i)
		break
	}
}
func (r *Redis) StartCleanInterval() {

	ticker := time.NewTicker(2 * time.Second)
	// quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				r.removeExpired()
				// case <-quit:
				// 	ticker.Stop()
				// 	return
			}
		}
	}()
}

func PriorityQCache() {
	testBreak()
	cache := NewCache()
	// cache.StartCleanInterval()
	cache.set("test key", "test value", time.Duration(2))
	cache.set("test key1", "test value", time.Duration(2))
	cache.set("test key2", "test value", time.Duration(3))
	cache.set("test key3", "test value", time.Duration(4))
	res := cache.get("test key")
	fmt.Println("result:", res)
	fmt.Println("results:", cache.items)
	time.Sleep(3 * time.Second)
	exp := cache.get("test key")
	fmt.Println("EXP KEY: ", exp)
	cache.removeExpired()
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
	log.Printf("HEAP==> %+v", *h)
}

func (h *ExpirationHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	fmt.Println("BALANCE ==> ", *h)
	return item
}

func (h ExpirationHeap) Peek() *expirationItem {
	if len(h) == 0 {
		return nil
	}
	return h[0]
}
