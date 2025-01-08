package counter

import (
	"errors"
	"strings"
)

func CountBytes(content string) (int, error) {
	if content == "" {
		return 0, errors.New("conteúdo vazio")
	}

	return int(len(content)), nil
}

func CountLynes(content string) (int, error) {
	if content == "" {
		return 0, errors.New("conteúdo vazio")
	}

	lineCount := strings.Count(string(content), "\n")

	if len(content) > 0 && content[len(content)-1] != '\n' {
		lineCount++
	}

	return lineCount, nil
}
