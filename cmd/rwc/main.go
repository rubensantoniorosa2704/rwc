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

	var file io.Reader     // Declare a variable for the file reader
	var inputSource string // Variable to store input source for printing

	if opt.InputFile != "" {
		// Try to open the provided file
		f, err := os.Open(opt.InputFile)
		if err != nil {
			fmt.Printf("Error opening file: %v", err)
			return
		}
		defer f.Close()
		file = f
		inputSource = opt.InputFile // Name of the input file for printing
	} else {
		// Use stdin if no file is provided
		file = os.Stdin
		inputSource = "stdin" // Indicate that the input is from stdin
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

	counts := make(map[string]int)

	// Create a scanner to read the file or stdin
	scanner := bufio.NewScanner(file)

	// Process each line in the input
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		for key, strategy := range strategiesMap {
			counts[key] += strategy.Count(line)
		}
	}

	// Check for reading errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
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

	// Print the input source (filename or 'stdin')
	fmt.Printf(" %s\n", inputSource)
}
