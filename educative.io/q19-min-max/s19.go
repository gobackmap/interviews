package main

import (
	"fmt"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	x, y := -42, 53
	fmt.Println("got:", []int{x, y})
	fmt.Println("max:", Max(x, y))
	fmt.Println("min:", Min(x, y))
}
