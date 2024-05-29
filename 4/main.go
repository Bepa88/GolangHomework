package main

import (
	"fmt"
	"strings"
)

func findContainsString(text []string, searchString string) []string {
	var result []string
	for _, v := range text {
		if strings.Contains(v, searchString) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	var text []string
	var input string
	var searchString string
	fmt.Print("Введіть текст (закінчіть введення цифрою 0):")

	for {
		if input == "0" {
			break
		}
		fmt.Scan(&input)
		text = append(text, input)
	}

	fmt.Print("Введіть рядок для пошуку:")
	fmt.Scan(&searchString)

	for _, v := range findContainsString(text, searchString) {
		fmt.Println(v)
	}
}
