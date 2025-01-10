package fileutils

import (
	"bufio"
	"os"
)

type CharCount struct {
	File *os.File
}

func (cc *CharCount) CountChar() (int, error) {
	scanner := bufio.NewScanner(cc.File)

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
