package biter

import (
	"fmt"

	entity "github.com/ZAF07/go-basics/composition/living-entity"
)

// Fields are not exported because we want this properties to be private (so nobody from the outside can call them directly, hence reducing vulnerabilities and runtime bugs)
type biter struct { // Struct is also private, so it cant be used standalone in any case which may introduce bugs
	biteForce            int
	entity.ILivingEntity // 💡 Embedded anonymous field containing the entity struct. This makes the Nonbiter struct looks like it inherits from the entity stuct. But it does not! It only allows us to call properties and methods of the embedded struct like it belongs to the child struct
	// entity.Entity // 💡 Use this method if you want to explicitly use the entity.Entity struct. The above (👆) field, uses an interface instead. So that makes it more extensible as ANY struct that implements the ILivingEntity interface is allowed to be used
	// enti entity.ILivingEntity // 💡 To go a step further in not exposing methods/fields to clients who use this struct, we can set a field name with a lowercase name. Then it is up to this struct to expose methods that calls methods of the embedded struct
}

func NewBiter(name string, bf int) *biter {
	// Using 'value' struct literal. (Have to pass the parameters in order of the struct fields in declaration of struct type)
	return &biter{
		bf,
		entity.NewLivingentity(name, 15),
	}

	// Using 'field:value' struct literal to init and return a biter struct
	// return &biter{
	// 	biteForce: bf,
	// 	Entity:    entity.NewLivingEntity(name, 15), // Here because we are creating the embedded struct via sturct literals (without using the 'new' keyword) we have to initialise the embedding explicitly. (Anonymous field: 'entity.Entity' (in the above struct declaration) == 'Entity'. So the embedding struct has a field name 'Entity')
	// }
}

func (b *biter) Bite(e entity.ILivingEntity) {
	fmt.Printf("🚨 %s is biting %s\n", b.GetName(), e.GetName())
	e.DecreaseHealth(b.biteForce)
}
