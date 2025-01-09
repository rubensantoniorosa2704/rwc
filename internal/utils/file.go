package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}

	return file, nil
}

func ReadFile(file *os.File, ch chan string) {
	defer close(ch)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linha := scanner.Text()
		ch <- linha
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
}

func ReadFileInChunks(file *os.File, ch chan<- []byte) {
	defer close(ch)

	const bufferSize = 64 * 1024
	buffer := make([]byte, bufferSize)
	totalBytesRead := 0

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}

		if n > 0 {
			totalBytesRead += n
			ch <- append([]byte(nil), buffer[:n]...)
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Total bytes read: %d\n", totalBytesRead)
}
