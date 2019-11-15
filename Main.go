package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}

	for range a {
		fmt.Println("1")
	}
}
