package fileutils

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

func CountFileMetrics(file *os.File) (int, int, int, int, error) {
	scanner := bufio.NewScanner(file)
	byteCount := 0
	lineCount := 0
	wordCount := 0
	charCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		byteCount += len(line) + 1 // +1 for the newline character
		lineCount++

		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			wordCount++
		}

		charCount += utf8.RuneCountInString(line)
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, 0, err
	}

	return byteCount, lineCount, wordCount, charCount, nil
}
