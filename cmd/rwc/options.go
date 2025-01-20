package main

import "flag"

type Options struct {
	CountBytes bool
	CountLines bool
	CountWords bool
	CountChar  bool
	InputFile  string
}

func ParseFlags() Options {
	countBytesFlag := flag.Bool("c", false, "Count bytes per file")
	countLinesFlag := flag.Bool("l", false, "Count lines per file")
	countWordsFlag := flag.Bool("w", false, "Count words per file")
	countCharacters := flag.Bool("m", false, "Count characters per file")

	flag.Parse()

	// Check if there are additional command-line arguments and assign the first one to inputFile
	inputFile := ""
	if len(flag.Args()) > 0 {
		inputFile = flag.Args()[0]
	}

	// Set default flags if none are provided
	if !*countBytesFlag && !*countLinesFlag && !*countWordsFlag && !*countCharacters {
		*countBytesFlag = true
		*countLinesFlag = true
		*countWordsFlag = true
	}

	return Options{
		CountBytes: *countBytesFlag,
		CountLines: *countLinesFlag,
		CountWords: *countWordsFlag,
		CountChar:  *countCharacters,
		InputFile:  inputFile,
	}
}
