package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	strategies "github.com/rubensantoniorosa2704/rwc/internal/fileutils/strategies"
)

func main() {
	// Retrieve available flags
	opt := ParseFlags()

	var file io.Reader

	if opt.InputFile != "" {
		// Try to open the provided file
		file, err := os.Open(opt.InputFile)
		if err != nil {
			fmt.Printf("Error opening file: %v", err)
			return
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	// Initialize counters
	var lineCount int
	strategiesMap := make(map[string]strategies.CounterStrategy)

	// Map flags to their respective strategies
	// Associating flags with corresponding strategies
	if opt.CountBytes {
		strategiesMap["bytes"] = strategies.ByteCount{}
	}
	if opt.CountWords {
		strategiesMap["words"] = strategies.WordCount{}
	}
	if opt.CountChar {
		strategiesMap["chars"] = strategies.CharCount{}
	}
	if opt.CountLines {
		strategiesMap["lines"] = strategies.LineCount{}
	}

	scanner := bufio.NewScanner(file)
	counts := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		for key, strategy := range strategiesMap {
			counts[key] += strategy.Count(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the results
	if opt.CountBytes {
		fmt.Printf(" %d", counts["bytes"])
	}
	if opt.CountLines {
		fmt.Printf(" %d", lineCount)
	}
	if opt.CountWords {
		fmt.Printf(" %d", counts["words"])
	}
	if opt.CountChar {
		fmt.Printf(" %d", counts["chars"])
	}
	fmt.Printf(" %s\n", opt.InputFile)
}
