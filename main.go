package main

import (
	"fmt"
	concurrency "go-adv/1-concurrency"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := make([]int, 0, 10)

	go concurrency.CreateSlice(ch1)
	go concurrency.Sqr(ch1, ch2)

	for v := range ch2 {
		arr = append(arr, v)
	}

	for _, value := range arr {
		fmt.Print(" ", value)
	}
}
