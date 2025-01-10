package fileutils

import (
	"bufio"
	"os"
)

type LineCount struct {
	File *os.File
}

func (lc *LineCount) CountLines() (int, error) {
	scanner := bufio.NewScanner(lc.File)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}
