package fileutils

import (
	"bufio"
	"os"
	"strings"
)

type WordCount struct {
	File *os.File
}

func (wc *WordCount) CountWords() (int, error) {
	scanner := bufio.NewScanner(wc.File)

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
