package strategies

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCount struct{}

func (w WordCount) Count(ch <-chan string) (int, error) {
	input := <-ch
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	var wordCount int = 0

	for scanner.Scan() {
		wordCount++
	}

	err := scanner.Err()
	if err != nil {
		return 0, fmt.Errorf("error scanning words: %w", err)
	}

	return wordCount, nil
}
