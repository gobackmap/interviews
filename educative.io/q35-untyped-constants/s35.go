package main

import "fmt"

type Number int

func main() {
	const c = 1
	// var c = 1	//uncomment to see the compile error
	fmt.Printf("Type of c is \"%T\", c=%d\n", c, c)
	var n Number = c
	fmt.Printf("Type of n is \"%T\", \"var n Number = c\" leads to n=%d\n\n ", n, n)
	fmt.Println(
		"Remember: x is assignable to T if x is an untyped constant representable by" +
			" a value of type T.",
	)
}
