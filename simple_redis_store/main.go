package main

import "fmt"

/*
This is an experiment to create an in-memory store like redis using Go
*/

func main() {
	fmt.Println("hello")
}

// Create the struct that represents the in-memory store
// Create methods to expose Get, Set, Update and Delete features
// Create a TTL method to delete a specific key:value when the TTL expires
