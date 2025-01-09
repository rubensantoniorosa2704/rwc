package counter

import (
	"os"
	"strings"

	"github.com/rubensantoniorosa2704/rwc/internal/utils"
)

func CountBytes(file *os.File) (int, error) {

	ch := make(chan []byte)
	go utils.ReadFileInChunks(file, ch)

	byteCount := 0

	for chunk := range ch {
		byteCount += len(chunk)
	}

	return byteCount, nil
}

func CountLines(file *os.File) (int, error) {
	ch := make(chan string)
	go utils.ReadFile(file, ch)

	lineCount := 0

	for line := range ch {
		_ = line
		lineCount++
	}

	return lineCount, nil
}

func CountWords(file *os.File) (int, error) {
	ch := make(chan string)
	go utils.ReadFile(file, ch)

	wordCount := 0

	for line := range ch {
		words := strings.Fields(line)
		wordCount += len(words)
	}

	return wordCount, nil
}
