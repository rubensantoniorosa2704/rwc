package strategies

import (
	"bufio"
	"fmt"
	"strings"
)

type LineCount struct{}

func (w LineCount) Count(ch <-chan string) (int, error) {
	input := <-ch
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
