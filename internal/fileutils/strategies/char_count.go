package strategies

import (
	"bufio"
	"fmt"
	"strings"
)

type CharCount struct{}

func (w CharCount) Count(ch <-chan string) (int, error) {
	input := <-ch
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
