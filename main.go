package main

import (
	"fmt"

	entity "github.com/ZAF07/go-basics/composition/living-entity"
)

func main() {
	shark := entity.NewBiter("sharkboy", 2)
	human := entity.NewNonBiter("human guy")
	fmt.Println(shark)
	fmt.Println(human)

	shark.Bite(human)
	fmt.Println(human)
	humanName := human.GetName() // If both the embedded struct and the struct that embeds the struct has the same method, Go will usee the struct's method instead
	fmt.Println(humanName)
}
