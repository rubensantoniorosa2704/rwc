package strategies

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCount struct{}

func (b ByteCount) Count(ch <-chan string) (int, error) {
	var byteCount int

	for input := range ch {
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanBytes)

		for scanner.Scan() {
			byteCount++
		}

		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("error scanning bytes: %w", err)
		}
	}

	return byteCount, nil
}
