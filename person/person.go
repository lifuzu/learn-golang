package person


import (
    "fmt"
)


type Person struct {
    Name string
}

func (p *Person) Talk() {
    fmt.Println("HI, MY NAME IS", p.Name)
}

// Embedded Types - `Person`

type Android struct {
    Person
    Model string
}

