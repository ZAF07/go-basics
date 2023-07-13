package livingentity

type NonBiter struct {
	entity // Embedded anonymous field containing the entity struct. This makes the NonBiter struct looks like it inherits from the entity stuct. But it does not! It only allows us to call properties and methods of the embedded struct like it belongs to the child struct
}

// Notice that the entity struct also implements the GetName() method. If both embedded and child struct implements the same methods, then the child struct's method will take precedence over the embedded struct's method
// func (n *NonBiter) GetName() string {
// 	return "Haha"
// }

func NewNonBiter(name string) *NonBiter {
	return &NonBiter{
		entity{name: name, health: 10},
	}
}
