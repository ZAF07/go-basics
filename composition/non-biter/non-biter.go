package livingentity

import (
	entity "github.com/ZAF07/go-basics/composition/living-entity"
)

/*
Embedded Field:
In Go, when you embed a struct type within another struct type, you can optionally provide a field name for the embedded struct. This field name becomes a field in the outer struct, and you access the fields and methods of the embedded struct through this field.
With an embedded field, you explicitly specify the field name for the embedded struct. This can be useful for disambiguating fields with the same name in the embedded struct and the outer struct.


Embedded Anonymous Field:
An embedded anonymous field in Go is a struct type that is embedded within another struct type without specifying a field name. This results in the fields and methods of the embedded struct being directly accessible from the outer struct without using a field name.
With an embedded anonymous field, you don't explicitly specify a field name for the embedded struct. Instead, you access its fields and methods as if they were part of the outer struct directly.
*/

type nonBiter struct {
	entity.ILivingEntity // Embedded anonymous field containing the entity struct. This makes the NonBiter struct looks like it inherits from the entity stuct (Just because we can pass the embedding struct directly to methods that accept the embedded struct's interface. Eg: human := NewNonBiter() method(human). Instead of human.EmbeddedStructFieldName). But it does not! It only allows us to call properties and methods of the embedded struct like it belongs to the child struct

	// Entity entity.ILivingEntity // Embedded field containing the entity struct. If the embedded struct implements an interface and we need to pass that interface type to a method, we have to do: human := NewNonBiter() method(human.Entity).
}

/*
	Notice that the NonBiter struct also implements the GetName() method.
	If both embedded and embedding struct implements the same methods, then the embedding struct's method will take precedence over the embedded struct's method
*/
// func (n *NonBiter) GetName() string {
// 	return "Haha"
// }

func NewNonBiter(name string) *nonBiter {
	return &nonBiter{
		entity.NewLivingentity(name, 10),
	}
}
