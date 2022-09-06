package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 10

	fmt.Println("The a variable is initialized to", a)
	fmt.Println("for loop starts")
	/* for loop execution */
	for a < 20 {
		fmt.Printf("  > value of a: %d\n", a)
		a++
		if a > 15 {
			/* terminate the loop using break statement */
			break
		}
	}
	fmt.Println("for loop is breaked!")
}
