package fileutils

import (
	"bufio"
	"fmt"
	"strings"
)

func CountBytes(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanBytes)
	var byteCount int = 0

	for scanner.Scan() {
		byteCount++
	}

	err := scanner.Err()
	if err != nil {
		return 0, fmt.Errorf("error scanning bytes: %w", err)
	}

	return byteCount, nil
}
