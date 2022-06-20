package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Printf("Hello, %s. You are at %s\n", os.Args[1], os.Args[0])
}
