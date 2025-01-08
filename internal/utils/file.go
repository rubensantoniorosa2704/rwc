package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo: %v", err)
	}

	if len(content) == 0 {
		return "", fmt.Errorf("o arquivo '%s' está vazio", file.Name())
	}

	return string(content), nil
}
