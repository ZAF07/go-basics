package livingentity

// Common interface for any living creature
// This allows polymorphism. Interfaces allows one struct to implement multiple interfaces. Allowing the Bite() method of the Biter struct to take in both a Biter and a NonBiter type because both types satisfies the ILivingCreature interface
type ILivingEntity interface {
	GetName() string
	GetHealth() int
	DecreaseHealth(p int) // Private method, not exported
}
