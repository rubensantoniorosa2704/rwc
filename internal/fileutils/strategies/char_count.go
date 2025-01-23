package strategies

import (
	"bufio"
	"fmt"
	"strings"
)

type CharCount struct{}

func (w CharCount) Count(ch <-chan string) (int, error) {
	var charCount int

	for input := range ch { // Itera sobre todas as entradas enviadas ao canal
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanRunes)

		for scanner.Scan() {
			charCount++
		}

		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("error scanning char: %w", err)
		}
	}

	return charCount, nil
}
