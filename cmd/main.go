package main

import (
	"fmt"
	"log"

	"github.com/rubensantoniorosa2704/rwc/internal/counter"
	"github.com/rubensantoniorosa2704/rwc/internal/utils"
)

func main() {
	// Parse flags
	opt := counter.ParseFlags()

	// Retrive the file content
	content, _ := utils.ReadFile(opt.InputFile)

	if opt.CountBytes {
		bytes, err := counter.CountBytes(content)
		if err != nil {
			log.Fatalf("Erro ao contar os bytes: %v\n", err)
		}
		fmt.Printf("Número de bytes: %d\n", bytes)
	}

	if opt.CountLines {
		lines, err := counter.CountLynes(content)
		if err != nil {
			log.Fatalf("Erro ao contar as linhas: %v\n", err)
		}
		fmt.Printf("Número de linhas: %d\n", lines)
	}
}
