package main

import (
	"container/heap"
	"fmt"

	q "github.com/ZAF07/go-basics/priority-queue"
)

// entity "github.com/ZAF07/go-basics/composition/living-entity"
// singleton "github.com/ZAF07/go-basics/sync/sync-once"

func main() {

	// // SYNC.ONCE
	// // Call the NewPewrson function. This causes sync.Once to register that the NewPerson has ran ONE TIME. It will NOT run again
	// p1 := singleton.NewPerson()
	// // We can prove that by changing the name value of the person singleton and call NewPerson() again. We will see that we are in fact acting on the same instance that was returned when calling NewPerson() the first time
	// p1.Name = "James"
	// fmt.Println(p1)

	// // Calling the second time will NOT execute the inner function of the NewPerson(). We can see that the 'NewPerson Ran' was only printed once
	// p2 := singleton.NewPerson()
	// fmt.Println(p2)

	// PQ (unrelated to concurrency)
	w := q.NewPQ()
	fmt.Println("Before push: ", w)
	a := q.PriorityItem{
		Value:   3,
		StringV: "3",
	}

	b := q.PriorityItem{
		Value:   1,
		StringV: "11",
	}

	c := q.PriorityItem{
		Value:   4,
		StringV: "4",
	}

	d := q.PriorityItem{
		Value:   2,
		StringV: "66",
	}

	heap.Push(w, &a)
	heap.Push(w, &b)
	heap.Push(w, &c)
	heap.Push(w, &d)

	fmt.Printf("after push: %+v\n", w)
	fmt.Println("peek --> ", w.Peek())
	ap := w.Pop()
	bp := w.Pop()
	cp := w.Pop()
	dp := w.Pop()

	fmt.Println("❌", ap)
	fmt.Println("❌", bp)
	fmt.Println("❌", cp)
	fmt.Println("❌", dp)
}
