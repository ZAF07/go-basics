package main

import (
	"fmt"

	"github.com/ZAF07/go-basics/composition"
)

func main() {
	shark := composition.NewBiter("sharkboy", 2)
	human := composition.NewNonBiter("human guy")
	fmt.Println(shark)
	fmt.Println(human)

	// shark.Bite(human)
	fmt.Println(human)
	a := human.GetName() // If both the embedded struct and the struct that embeds the struct has the same method, Go will usee the struct's method instead
	fmt.Println(a)
}
