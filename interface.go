package main

import "fmt"

type HashName interface {
	GetName() string
}

func SayHello(value HashName) {
	fmt.Println("Hello", value.GetName())
}

// implementasi interface
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// end implementasi interface
func main() {
	person := Person{Name: "Ricko"} //set person
	SayHello(person)                //call method
	animal := Animal{Name: "Gajah"} //set person
	SayHello(animal)                //call method
}
