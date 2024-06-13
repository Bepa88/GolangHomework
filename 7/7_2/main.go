package main

import (
	"fmt"
	"math/rand"
)

func main() {

	intSliceCh := make(chan []int)
	intChMin := make(chan int)
	intChMax := make(chan int)

	go rendom(100, intSliceCh, intChMin, intChMax)
	go minMax(<-intSliceCh, intChMin, intChMax)
	fmt.Scanln()
}

func rendom(n int, ch chan []int, min chan int, max chan int) {
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, rand.Intn(100))
	}
	ch <- result
	close(ch)
	fmt.Println(<-min)
	fmt.Println(<-max)
}

func minMax(randomSlice []int, min chan int, max chan int) {
	maxVal := randomSlice[0]
	minVal := randomSlice[0]

	// Прохід через всі елементи масиву
	for _, value := range randomSlice {
		if value > maxVal {
			maxVal = value
		}
		if value < minVal {
			minVal = value
		}
	}
	min <- maxVal
	max <- minVal
	close(min)
	close(max)
}
