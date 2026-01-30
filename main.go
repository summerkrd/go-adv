package main

import (
	random_api "go-adv/2-random-api"
	"net/http"
)

func main() {
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//arr := make([]int, 0, 10)
	//
	//go concurrency.CreateSlice(ch1)
	//go concurrency.Sqr(ch1, ch2)
	//
	//for v := range ch2 {
	//	arr = append(arr, v)
	//}
	//
	//for _, value := range arr {
	//	fmt.Print(" ", value)
	//}

	router := http.NewServeMux()
	randApi := &random_api.RandomApi{}
	randApi.StartApi(router)
}
