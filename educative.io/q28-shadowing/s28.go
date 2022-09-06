package main

import (
	"fmt"
)

func main() {
	x := 0

	// Print the value
	fmt.Println("Before the decision block, x:", x)

	toBeDeclared := false

	if toBeDeclared { // x:= 1 declaration shadows the original x
		x := 1
		x++
	}

	if !toBeDeclared { // want to reuse the actual var x from the outer block?
		x = 2
		x++
	}

	// Print the value
	fmt.Println("After the decision block, x:", x)
}
