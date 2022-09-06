package main

import (
	"fmt"
)

func ckeckIfEmpty(slice []any) string {
	switch len(slice) {
	case 0:
		return "empty!"
	default:
		return "not empty!"
	}
}

func main() {
	slice := []any{"string", 12}
	fmt.Println("got", slice)
	fmt.Println("This slice is", ckeckIfEmpty(slice))
}
