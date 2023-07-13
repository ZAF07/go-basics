package livingentity

// Implements the ILivingEntity interface
// Fields are not exported because we want this propertues to be private (so nobody from the outside can call them directly, hence less vulnerabilities of bugs)
type entity struct {
	name   string
	health int
}

func NewLivingEntity(name string, health int) *entity {
	return &entity{
		name:   name,
		health: health,
	}
}

func (e *entity) GetName() string {
	return e.name
}

// Fields are not exported because we dont want the decreaseHealth method to be public
func (e *entity) decreaseHealth(p int) {
	e.health -= p
}
