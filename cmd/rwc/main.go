package main

import (
	"fmt"
	"os"

	fileutils "github.com/rubensantoniorosa2704/rwc/internal/fileutils/src"
)

func main() {
	opt := ParseFlags()

	file, err := os.Open(opt.InputFile)
	if err != nil {
		fmt.Printf("erro ao abrir o arquivo: %v", err)
	}

	switch {
	case opt.CountBytes:
		byteCount, err := fileutils.CountBytes(file)
		if err != nil {
			fmt.Printf("Erro ao contar os bytes: %v\n", err)
		}
		fmt.Printf("Número de bytes: %d\n", byteCount)

	case opt.CountLines:
		lineCount, err := fileutils.CountLines(file)
		if err != nil {
			fmt.Printf("Erro ao contar as linhas: %v\n", err)
		}
		fmt.Printf("Número de linhas: %d\n", lineCount)

	case opt.CountWords:
		wordCount, err := fileutils.CountWords(file)
		if err != nil {
			fmt.Printf("Erro ao contar as palavras: %v\n", err)
		}
		fmt.Printf("Número de palavras: %d\n", wordCount)

	case opt.CountChar:
		characterCount, err := fileutils.CountChar(file)
		if err != nil {
			fmt.Printf("Erro ao contar os caracteres: %v\n", err)
		}
		fmt.Printf("Número de caracteres: %d\n", characterCount)
	}
}
