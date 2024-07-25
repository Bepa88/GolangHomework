package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func main() {
	const filename = "1689007675141_numbers.txt"
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}

	re := regexp.MustCompile(`\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)
	matches := re.FindAllString(string(fileContent), -1)

	fmt.Println("Found phone numbers:")
	for _, match := range matches {
		fmt.Println(match)
	}

}
