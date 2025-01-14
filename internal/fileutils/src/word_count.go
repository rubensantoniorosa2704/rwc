package fileutils

import (
	"bufio"
	"os"
	"strings"
)

func CountWords(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)

	wordCount := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}
