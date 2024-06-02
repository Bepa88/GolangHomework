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
	j := 0
	for i := 1; i < len(slice); i++ {
		if slice[j].ID != slice[i].ID {
			j++
			slice[j] = slice[i]
		}
	}
	return slice[:j+1]
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
