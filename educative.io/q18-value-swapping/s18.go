package main

import "fmt"

func swap(x, y int) []int {
	y, x = x, y
	return []int{x, y}
}
func main() {
	a, b := 5, 10
	fmt.Println("got:", []int{a, b})
	fmt.Println("swapped:", swap(a, b))
}
