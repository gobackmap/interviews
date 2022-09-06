package main

import (
	"fmt"
	"os"
)

// Golang program to show how to
// use command-line arguments.
func main() {

	// The first argument is always program name
	myProgramName := os.Args[0]

	// this will take 4th command line arguments
	cmdArgs := os.Args[4]

	// getting the arguments with normal indexing
	gettingArgs := os.Args[2]

	toGetAllArgs := os.Args[1:]

	// it will display the program name
	fmt.Println("Program name:", myProgramName)
	fmt.Println("cmdArgs:", cmdArgs)
	fmt.Println("gettingArgs:", gettingArgs)
	fmt.Println("toGetAllArgs:", toGetAllArgs)
}

// Panic?! Call the program with this command:

// go run q40-cl_arguments/s40.go 1st 2nd 3d 4th
