package wordlist

import (
	"bufio"
	"fmt"
	"os"
)

func ReadWordlistFile(wordlistPath string) []string {
	file, err := os.Open(wordlistPath)

	if err != nil {
		fmt.Println("Wordlist error ->", err)
	}

	scanner := bufio.NewScanner(file)
	wordArray := []string{}
	for scanner.Scan() {
		wordArray = append(wordArray, scanner.Text())
	}

	return wordArray
}
