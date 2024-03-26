package main

import (
	"fmt"

	"github.com/ZAF07/go-basics/composition/biter"
	nonBiter "github.com/ZAF07/go-basics/composition/non-biter"
)

func main() {
	// Init new biter and non-biter
	shark := biter.NewBiter("Shark", 3)
	human := nonBiter.NewNonBiter("Human")

	// Print the initial health of both entities
	fmt.Println("💡 Shark health: ", shark.GetHealth())
	fmt.Println("💡 Human health: ", human.GetHealth())

	// Get biter to bite a nonBiter
	shark.Bite(human)

	// Print the health of the entity the biter bit
	fmt.Println("💡 Human health after being bitten: ", human.GetHealth())

	// Init a new biter to show that a biter can also bite a biter because they are also living entities
	bittenBiter := biter.NewBiter("Giant Squid", 1)
	fmt.Println("💡 bittenBiter's health before being bitten: ", bittenBiter.GetHealth())
	// Get biter to bite the bittenBiter
	shark.Bite(bittenBiter)

	// Now print the health of the bittenBiter after being bitten
	fmt.Println("💡 Health of bittenBiter after being bitten: ", bittenBiter.GetHealth())
}
