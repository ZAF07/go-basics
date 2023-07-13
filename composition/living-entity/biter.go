package livingentity

// Fields are not exported because we want this properties to be private (so nobody from the outside can call them directly, hence reduced vulnerabilities and bugs)
type Biter struct {
	biteForce int
	entity    // Embedded anonymous field containing the entity struct. This makes the NonBiter struct looks like it inherits from the entity stuct. But it does not! It only allows us to call properties and methods of the embedded struct like it belongs to the child struct
}

func NewBiter(name string, bf int) *Biter {
	return &Biter{
		biteForce: bf,
		entity:    entity{name: name, health: 15},
	}
}

func (b *Biter) Bite(e ILivingEntity) {
	e.decreaseHealth(b.biteForce)
}
