package main

import "fmt"

func main() {
	// Create
	var stack []string

	// Push
	stack = append(stack, "World!")
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		// Print top
		n := len(stack) - 1
		fmt.Print(stack[n])

		// Pop
		stack = stack[:n]
	}
	fmt.Println("")
	// Output : Hello World!
}
