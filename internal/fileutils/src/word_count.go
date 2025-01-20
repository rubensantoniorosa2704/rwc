package fileutils

import (
	"bufio"
	"fmt"
	"strings"
)

func CountWords(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	var wordCount int = 0

	for scanner.Scan() {
		wordCount++
	}

	err := scanner.Err()
	if err != nil {
		return 0, fmt.Errorf("error scanning words: %w", err)
	}

	return wordCount, nil
}
