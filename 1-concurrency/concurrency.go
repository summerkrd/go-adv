package concurrency

import (
	"math/rand"
	"time"
)

func CreateSlice(ch chan int) {
	randSlice := make([]int, 0, 10)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		randSlice = append(randSlice, rnd.Intn(101))
	}

	for _, value := range randSlice {
		ch <- value
	}

	close(ch)
}

func Sqr(inputCh <-chan int, outputCh chan<- int) {
	for value := range inputCh {
		outputCh <- value * value
	}
	close(outputCh)
}
