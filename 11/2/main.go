package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func main() {
	const filename = "1689007676028_text.txt"
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}

	vowelConsonantPattern := `(?:\s|^|[^\p{L}])([аеєиіїоуюяАЕЄИІЇОУЮЯ][а-яА-Я]*[бвгґджзклмнпрстфхцчшщБВГҐДЖЗКЛМНПРСТФХЦЧШЩ])(?:\s|$|[^\p{L}])`
	reVowelConsonant := regexp.MustCompile(vowelConsonantPattern)
	matchesVowelConsonant := reVowelConsonant.FindAllString(string(fileContent), -1)

	fmt.Println("Слова, що починаються на голосні та закінчуються на приголосні:")
	for _, match := range matchesVowelConsonant {
		fmt.Println(match)
	}
}
