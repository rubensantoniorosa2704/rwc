package strategies

import (
	"bufio"
	"fmt"
	"strings"
)

type LineCount struct{}

func (w LineCount) Count(ch <-chan string) (int, error) {
	var lineCount int

	for input := range ch { // Itera sobre todas as entradas enviadas ao canal
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			lineCount++
		}

		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("error scanning lines: %w", err)
		}
	}

	return lineCount, nil
}
