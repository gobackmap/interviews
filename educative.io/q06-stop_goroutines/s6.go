package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	fmt.Println("Starting the go routine")
	i := 1
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Printf("go routine number %d is done!\n", i)
				i++
			}
		}
	}()
	time.Sleep(20 * time.Microsecond)
	quit <- true
	fmt.Println("Go routine is stopped")
}
