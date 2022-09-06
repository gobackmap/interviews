package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {

		if len(left) == 0 {
			return append(merged, right...)

		} else if len(right) == 0 {
			return append(merged, left...)

		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]

		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func SequentialMergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	left := SequentialMergeSort(data[:mid])
	right := SequentialMergeSort(data[mid:])
	return Merge(left, right)
}

func ConcurrentMergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	done := make(chan bool)
	var left []int
	go func() {
		left = ConcurrentMergeSort(data[:mid])
		done <- true
	}()
	right := ConcurrentMergeSort(data[mid:])
	<-done
	return Merge(left, right)
}

func main() {
	// data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	rand.Seed(time.Now().Unix())
	data := rand.Perm(20)
	fmt.Println("data:", data)
	defer track("SequentialMergeSort")()
	fmt.Println("sequential:", SequentialMergeSort(data))
	defer track("ConcurrentMergeSort")()
	fmt.Println("concurrently:", ConcurrentMergeSort(data))
}

func track(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s, execution time %s\n", name, time.Since(start))
	}
}
