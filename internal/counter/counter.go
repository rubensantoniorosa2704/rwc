package counter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/rubensantoniorosa2704/rwc/internal/utils"
)

func CountBytes(file *os.File) (int, error) {

	if file == nil {
		return 0, fmt.Errorf("arquivo inválido")
	}

	ch := make(chan []byte)

	go utils.ReadFileInChunks(file, ch)

	totalBytes := 0

	for chunk := range ch {
		scanner := bufio.NewScanner(bytes.NewReader(chunk))
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			totalBytes++
		}

		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("erro ao processar os dados: %w", err)
		}
	}

	return totalBytes, nil
}

func CountLines(file *os.File) (int, error) {
	if file == nil {
		return 0, fmt.Errorf("arquivo inválido")
	}

	ch := make(chan string)

	go utils.ReadFile(file, ch)

	lineCount := 0

	for range ch {
		lineCount++
	}

	return lineCount, nil
}

func CountWords(file *os.File) (int, error) {
	if file == nil {
		return 0, fmt.Errorf("arquivo inválido")
	}

	ch := make(chan string)

	go utils.ReadFile(file, ch)

	wordCount := 0

	for linha := range ch {
		words := strings.Fields(linha)
		wordCount += len(words)
	}

	return wordCount, nil
}

func CountCharacters(file *os.File) (int, error) {
	var totalCharacters int

	const bufferSize = 64 * 1024
	buffer := make([]byte, bufferSize)

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return 0, fmt.Errorf("erro ao ler o arquivo: %v", err)
		}

		for i := 0; i < n; {
			_, size := utf8.DecodeRune(buffer[i:])
			if size == 0 {
				return 0, fmt.Errorf("erro ao decodificar rune")
			}
			totalCharacters++
			i += size
		}

		if err == io.EOF {
			break
		}
	}

	return totalCharacters, nil
}
