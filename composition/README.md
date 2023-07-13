## Composition in Go

Go does not allow Inheritance, however, it does support Composition. Composition means that an object is composed of one or multiple objects, establishing a "HAS A" relationship instead of an "IS A" relationship found in Inheritance.

For example, a car **HAS A** wheel, a car **HAS A** windshield, and a car **HAS A** door.

In Go, this can be achieved using the concept of enbedded anonymous field or named Struct field method.

### Embedding (Embedded Anonymous Field)

Embedding a Struct in another Struct in Go allows us to call methods/properties of the embedded Struct (Parent Struct) as if they belonged to the Child Struct.

```go
type Wheel struct {
    // Wheel properties...
}

type Car struct {
    Wheel // Embedding the Wheel struct
    // Other Car properties...
}

func main() {
    c := Car{}
    c.Wheel.Method() // Accessing the Method of the embedded Wheel struct
}
```

In the example above, the Car struct embeds the Wheel struct. This enables us to call methods or access properties of the Wheel struct directly from an instance of the Car struct.

This composition approach allows us to create complex structures by combining multiple smaller components or objects. It promotes code reuse and modularity by encapsulating related functionalities within separate structs.

Remember that the embedded struct becomes a part of the parent struct and inherits its properties and methods. However, it's important to note that embedding is not inheritance, as Go does not have native support for classical inheritance.

By leveraging composition through embedding, we can achieve flexible and modular designs in Go.