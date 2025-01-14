package fileutils

import (
	"bufio"
	"os"
)

func CountLines(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}
