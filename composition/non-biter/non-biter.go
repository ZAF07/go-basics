package livingentity

import (
	entity "github.com/ZAF07/go-basics/composition/living-entity"
)

type NonBiter struct {
	entity.ILivingEntity // Embedded anonymous field containing the entity struct. This makes the NonBiter struct looks like it inherits from the entity stuct. But it does not! It only allows us to call properties and methods of the embedded struct like it belongs to the child struct
}

/*
	Notice that the NonBiter struct also implements the GetName() method.
	If both embedded and embedding struct implements the same methods, then the embedding struct's method will take precedence over the embedded struct's method
*/
// func (n *NonBiter) GetName() string {
// 	return "Haha"
// }

func NewNonBiter(name string) *NonBiter {
	return &NonBiter{
		entity.NewLivingentity(name, 10),
	}
}
