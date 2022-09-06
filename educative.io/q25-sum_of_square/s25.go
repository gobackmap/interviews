package main

import "fmt"

func SumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}

func main() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 0; i < 5; i++ {
			sum += <-mychannel
			fmt.Printf("i = %d, sum = %d\n", i, sum)
		}
		fmt.Println("=====> sum =", sum)
		quitchannel <- 0
	}()
	SumOfSquares(mychannel, quitchannel)
}
