package main

import (
	"fmt"
	"strings"
)

/*
func main() {
	// arrays
	data := [4]int{1, 2, 3, 4}

	for index, value := range data {
		println(index, value)
	}

	scores := make([]int, 0, 10)
	scores = scores[0:8]
	scores[7] = 90

	fmt.Println(scores)
}*/

/*
func main() {
	// slices
	scores := make([]int, 0, 5)

	c := cap(scores)

	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
}*/

func main() {
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4]
	slice[0] = 999
	fmt.Println(scores)
	fmt.Println(slice)

	haystacks := "the spice must flow"
	haystacks_slice := haystacks[:len(haystacks)-1]
	aha := strings.Index(haystacks_slice, " ")
	fmt.Println(haystacks_slice)
	fmt.Println(aha)
}
