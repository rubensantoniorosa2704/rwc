package strategies

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCount struct{}

func (w WordCount) Count(ch <-chan string) (int, error) {
	var wordCount int = 0

	for input := range ch { // Itera sobre todas as linhas enviadas ao canal
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			wordCount++
		}

		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("error scanning words: %w", err)
		}
	}

	return wordCount, nil
}
