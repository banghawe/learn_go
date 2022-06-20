package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func (s *Saiyan) Super() {
	s.Power += 5000
}

func main() {
	goku := &Saiyan{
		Name:  "goku",
		Power: 6000,
	}

	gohan := &Saiyan{
		Name:   "gohan",
		Power:  4000,
		Father: goku,
	}

	fmt.Printf("%v address is %p\n", goku.Name, goku)
	// super(goku)
	fmt.Printf("%s power is %d\n", goku.Name, goku.Power)
	goku.Super()
	fmt.Printf("%s power is %d\n", goku.Name, goku.Power)
	fmt.Printf("%s's father is %s\n", gohan.Name, gohan.Father.Name)
}

func super(s *Saiyan) {
	s.Power += 10000
	fmt.Printf("%p\n", s)
}
