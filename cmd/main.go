package main

import (
	"fmt"
	"log"

	"github.com/rubensantoniorosa2704/rwc/internal/counter"
	"github.com/rubensantoniorosa2704/rwc/internal/utils"
)

func main() {
	opt := counter.ParseFlags()

	file, err := utils.OpenFile(opt.InputFile)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v\n", err)
	}

	switch {
	case opt.CountBytes:
		byteCount, err := counter.CountBytes(file)
		if err != nil {
			log.Fatalf("Erro ao contar os bytes: %v\n", err)
		}
		fmt.Printf("Número de bytes: %d\n", byteCount)

	case opt.CountLines:
		lineCount, err := counter.CountLines(file)
		if err != nil {
			log.Fatalf("Erro ao contar as linhas: %v\n", err)
		}
		fmt.Printf("Número de linhas: %d\n", lineCount)

	case opt.CountWords:
		wordCount, err := counter.CountWords(file)
		if err != nil {
			log.Fatalf("Erro ao contar as palavras: %v\n", err)
		}
		fmt.Printf("Número de palavras: %d\n", wordCount)
	}
}
