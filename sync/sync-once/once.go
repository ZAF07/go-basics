package synconce

import (
	"fmt"
	"sync"
)

/* Sync Once
Package sync.Once in Go is a package that ensures a function executes EXACTLY ONCE no matter how many times it is called

In this example, we are going to implement the Singleton pattern in Go. This is a perfect example to demonstrate how the sync.Once package works
*/

// The singleton object
var Once sync.Once

var person *Person

type Person struct {
	Name string
	Age  int
}

func NewPerson() *Person {
	Once.Do(func() {
		fmt.Println("NewPerson Ran")
		p := &Person{
			Name: "Default person",
			Age:  100,
		}
		person = p
	})

	return person
}
