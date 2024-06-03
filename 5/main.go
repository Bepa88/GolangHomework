package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var note = make(map[string][]string)

func indexTextByWords(text []string) {

	for _, v := range text {
		words := strings.Fields(v)
		for _, word := range words {
			note[word] = append(note[word], v)
		}
	}
}

func findLinesByWord(word string) []string {
	if lines, ok := note[word]; ok {
		return lines
	}
	return []string{"Немає рядка зі словом: " + word}
}

func main() {
	var text []string
	var searchString string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введіть текст (закінчіть введення цифрою 0):")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "0" {
			break
		}
		text = append(text, input)
	}
	fmt.Print("Введіть рядок для пошуку:")
	fmt.Scan(&searchString)
	indexTextByWords(text)

	lines := findLinesByWord(searchString)
	for _, line := range lines {
		fmt.Println(line)
	}
}
