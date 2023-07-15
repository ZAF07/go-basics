package main

import (
	"fmt"

	entity "github.com/ZAF07/go-basics/composition/living-entity"
	singleton "github.com/ZAF07/go-basics/sync/sync-once"
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

	// SYNC.ONCE
	// Call the NewPewrson function. This causes sync.Once to register that the NewPerson has ran ONE TIME. It will NOT run again
	p1 := singleton.NewPerson()
	// We can prove that by changing the name value of the person singleton and call NewPerson() again. We will see that we are in fact acting on the same instance that was returned when calling NewPerson() the first time
	p1.Name = "James"
	fmt.Println(p1)

	// Calling the second time will NOT execute the inner function of the NewPerson(). We can see that the 'NewPerson Ran' was only printed once
	p2 := singleton.NewPerson()
	fmt.Println(p2)
}
