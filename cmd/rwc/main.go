package main

import (
	"fmt"
	"os"

	fileutils "github.com/rubensantoniorosa2704/rwc/internal/fileutils/src"
)

func main() {
	// Retrieve available flags
	opt := ParseFlags()

	// Try to open the provided file
	file, err := os.Open(opt.InputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	byteCount, lineCount, wordCount, charCount, err := fileutils.CountFileMetrics(file)
	if err != nil {
		fmt.Printf("Error counting metrics: %v\n", err)
		return
	}

	switch {
	case opt.CountBytes:
		fmt.Printf("Number of bytes: %d\n", byteCount)
	case opt.CountLines:
		fmt.Printf("Number of lines: %d\n", lineCount)
	case opt.CountWords:
		fmt.Printf("Number of words: %d\n", wordCount)
	case opt.CountChar:
		fmt.Printf("Number of characters: %d\n", charCount)
	default:
		fmt.Printf("%d %d %d\n", byteCount, lineCount, wordCount)
	}
}
