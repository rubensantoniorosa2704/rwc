package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	strategies "github.com/rubensantoniorosa2704/rwc/internal/fileutils/strategies"
)

func main() {
	opt := ParseFlags()

	var file io.Reader
	var inputSource string

	if opt.InputFile != "" {
		f, err := os.Open(opt.InputFile)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()
		file = f
		inputSource = opt.InputFile
	} else {
		file = os.Stdin
		inputSource = "stdin"
	}

	ch := make(chan string, 100)
	var wg sync.WaitGroup
	var mu sync.Mutex

	strategiesMap := make(map[string]strategies.CounterStrategy)

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

	for key, strategy := range strategiesMap {
		wg.Add(1)
		go func(key string, strategy strategies.CounterStrategy) {
			defer wg.Done()
			count, err := strategy.Count(ch)
			if err != nil {
				fmt.Printf("Error counting %s: %v\n", key, err)
				return
			}
			mu.Lock()
			counts[key] = count
			mu.Unlock()
		}(key, strategy)
	}

	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	if opt.CountBytes {
		fmt.Printf(" %d", counts["bytes"])
	}
	if opt.CountLines {
		fmt.Printf(" %d", counts["lines"])
	}
	if opt.CountWords {
		fmt.Printf(" %d", counts["words"])
	}
	if opt.CountChar {
		fmt.Printf(" %d", counts["chars"])
	}

	fmt.Printf(" %s\n", inputSource)
}
