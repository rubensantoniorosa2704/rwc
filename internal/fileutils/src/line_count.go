package fileutils

import (
	"bufio"
	"fmt"
	"strings"
)

func CountLines(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	var lineCount int = 0

	for scanner.Scan() {
		lineCount++
	}

	err := scanner.Err()
	if err != nil {
		return 0, fmt.Errorf("error scanning lines: %w", err)
	}

	return lineCount, nil
}
