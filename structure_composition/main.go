package main

import "fmt"

type Person struct {
	Name string
}

func (s *Person) Introduce() {
	fmt.Printf("My name is %s\n", s.Name)
}

type Saiyan struct {
	*Person
	Power int
}

// Overlaoding
func (s *Saiyan) Introduce() {
	fmt.Printf("I'm %s\n", s.Name)
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  6000,
	}
	goku.Introduce()
	goku.Person.Introduce()
}
