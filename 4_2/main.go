package main

import (
	"fmt"
	"sort"
)

type MyStruct struct {
	ID string
}

func UniqueData(slice []MyStruct) []MyStruct {
	sort.Slice(slice, func(i, j int) bool { return slice[i].ID < slice[j].ID })
	for i := 1; i < len(slice); i++ {
		if slice[i].ID == slice[i-1].ID {
			slice = append(slice[:i-1], slice[i:]...)
			i--
		}
	}
	return slice
}

func main() {
	var input string
	var slice []MyStruct
	fmt.Print("Введіть ID (закінчіть введення f):")

	for {
		fmt.Scan(&input)
		if input == "f" {
			break
		}
		str := MyStruct{
			ID: input,
		}
		slice = append(slice, str)
	}

	for _, v := range UniqueData(slice) {
		fmt.Println(v)
	}
}
