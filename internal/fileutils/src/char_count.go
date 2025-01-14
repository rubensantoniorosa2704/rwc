package fileutils

import (
	"bufio"
	"os"
)

func CountChar(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)

	charCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		charCount += len(line)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return charCount, nil
}
