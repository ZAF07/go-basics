package main

import (
	"container/heap"
	"fmt"
)

// 💡 This is a min PQ because in the Less(), we are comparing i < j
type PriorityItem struct {
	Value    interface{}
	Priority int
	Index    int
	StringV  string
}

type MinPriorityQ []*PriorityItem

func NewMinPQ() *MinPriorityQ {
	// q := make(MinPriorityQ, 0)
	q := &MinPriorityQ{}
	heap.Init(q)
	return q
}

func (pq MinPriorityQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]

}

// TODO: Learn what this does. Also look into heap.Fix()
func (pq MinPriorityQ) Update(item *PriorityItem, value interface{}, priority int) {
	fmt.Println("🚨 UPDATED ---> ", value, priority)
	item.Value = value
	item.Priority = priority
	heap.Fix(&pq, item.Index)
}

func (pq *MinPriorityQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *MinPriorityQ) Push(x interface{}) {
	*pq = append(*pq, x.(*PriorityItem))
}

// Less takes the index of the items to compare. Logic goes here
func (pq MinPriorityQ) Less(i, j int) bool {
	return pq[i].Value.(int) < pq[j].Value.(int)
}

func (pq MinPriorityQ) Len() int {
	return len(pq)
}

func (pq MinPriorityQ) Peek() interface{} {
	if len(pq) < 1 {
		return nil
	}
	return pq[len(pq)-1]
}

func main() {
	w := NewMinPQ()
	fmt.Println("Length before push: ", w.Len())
	a := PriorityItem{
		Value:   3,
		StringV: "3",
	}

	b := PriorityItem{
		Value:   1,
		StringV: "1",
	}

	c := PriorityItem{
		Value:   4,
		StringV: "4",
	}

	d := PriorityItem{
		Value:   2,
		StringV: "2",
	}

	heap.Push(w, &a)
	heap.Push(w, &c)
	heap.Push(w, &b)
	heap.Push(w, &d)

	fmt.Printf("Length after push: %+v\n", w.Len())
	fmt.Println("peek --> ", w.Peek())

	ap := heap.Pop(w)
	bp := heap.Pop(w)
	cp := heap.Pop(w)
	dp := heap.Pop(w)

	fmt.Println("✅", ap)
	fmt.Println("✅", bp)
	fmt.Println("✅", cp)
	fmt.Println("✅", dp)
	fmt.Println("Length after pop ---> ", w.Len())
}
