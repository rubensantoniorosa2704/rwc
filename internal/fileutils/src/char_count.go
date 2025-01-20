package fileutils

import (
	"bufio"
	"fmt"
	"strings"
)

func CountChar(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	var charCount int = 0

	for scanner.Scan() {
		charCount++
	}

	err := scanner.Err()
	if err != nil {
		return 0, fmt.Errorf("error scanning char: %w", err)
	}

	return charCount, nil
}
