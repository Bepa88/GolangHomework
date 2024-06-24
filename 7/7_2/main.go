package main

import (
	"fmt"
	"math/rand"
)

type MinMax struct {
	min int
	max int
}

func main() {

	intSliceCh := make(chan []int)
	chMinMax := make(chan MinMax)

	go rendom(100, intSliceCh, chMinMax)
	go minMax(intSliceCh, chMinMax)
	fmt.Scanln()
}

func rendom(n int, ch chan []int, minMax chan MinMax) {
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, rand.Intn(100))
	}
	ch <- result
	close(ch)
	minMaxV := <-minMax
	fmt.Println(minMaxV.min)
	fmt.Println(minMaxV.max)
}

func minMax(random chan []int, minMax chan MinMax) {
	randomSlice := <-random
	var minMaxVal MinMax
	for _, value := range randomSlice {
		if value > minMaxVal.max {
			minMaxVal.max = value
		}
		if value < minMaxVal.min {
			minMaxVal.min = value
		}
	}
	minMax <- minMaxVal
	close(minMax)
}
