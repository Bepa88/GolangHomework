package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intSliceCh := make(chan []int)
	intCh := make(chan int)

	go rendom(100, intSliceCh)
	go averageValue(intSliceCh, intCh)
	go printSum(intCh)
	fmt.Scanln()
}

func rendom(n int, ch chan []int) {
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, rand.Intn(100))
	}
	ch <- result
	close(ch)
}

func averageValue(random chan []int, ch chan int) {
	randomSlice := <-random
	capCh := len(randomSlice)
	var sum int
	for _, value := range randomSlice {
		sum += value
	}
	ch <- sum / capCh
	close(ch)
}

func printSum(sumCh chan int) {
	sum := <-sumCh
	fmt.Println(sum)
}
