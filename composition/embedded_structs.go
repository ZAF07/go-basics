package composition

// /* Composition in Go
// Go does not allow Inheritance, however it does support Composition
// Composition means that an Object is composed of one or multiple Objects (HAS A relationship as opposed to IS A relationship in Inheritance)

// Example: A car HAS A wheel. A car HAS A windshield. A car HAS A door...

// This means that in one Struct, (A struct in Go is somewhat similar to a Class in Java) we can have one or more other Structs.
// This is done with the embedding or Struct field method

// Embedding:
// Embedding a Struct in another Struct in Go allows us to call methods/properties of the embedded Struct (Parent Struct) as if it belonged to the Child Struct

// */

// // Common interface for any living creature
// // This allows polymorphism. Interfaces allows one struct to implement multiple interfaces. Allowing the Bite() method of the Biter struct to take in both a Biter and a NonBiter type because both types satisfies the ILivingCreature interface
// type ILivingCreature interface {
// 	GetName() string
// 	decreaseHealth(p int) // Private method, not exported
// }

// // Implements the interface
// // Fields are not exported because we want this propertues to be private (so nobody from the outside can call them directly, hence less vulnerabilities of bugs)
// type entity struct {
// 	name   string
// 	health int
// }

// func (e *entity) GetName() string {
// 	return e.name
// }

// // Fields are not exported because we dont want the decreaseHealth method to be public
// func (e *entity) decreaseHealth(p int) {
// 	e.health -= p
// }

// // Fields are not exported because we want this properties to be private (so nobody from the outside can call them directly, hence reduced vulnerabilities and bugs)
// type Biter struct {
// 	biteForce int
// 	entity    // Embedded anonymous field containing the entity struct. This makes the NonBiter struct looks like it inherits from the entity stuct. But it does not! It only allows us to call properties and methods of the embedded struct like it belongs to the child struct
// }

// func (b *Biter) Bite(e ILivingCreature) {
// 	e.decreaseHealth(b.biteForce)
// }

// func NewBiter(name string, bf int) *Biter {
// 	return &Biter{
// 		biteForce: bf,
// 		entity:    entity{name: name, health: 15},
// 	}
// }

// type NonBiter struct {
// 	entity // Embedded anonymous field containing the entity struct. This makes the NonBiter struct looks like it inherits from the entity stuct. But it does not! It only allows us to call properties and methods of the embedded struct like it belongs to the child struct
// }

// // Notice that the entity struct also implements the GetName() method. If both embedded and child struct implements the same methods, then the child struct's method will take precedence over the embedded struct's method
// // func (n *NonBiter) GetName() string {
// // 	return "Haha"
// // }

// func NewNonBiter(name string) *NonBiter {
// 	return &NonBiter{
// 		entity{name: name, health: 10},
// 	}
// }
