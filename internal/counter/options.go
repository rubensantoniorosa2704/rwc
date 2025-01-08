package counter

import "flag"

type Options struct {
	CountBytes bool
	CountLines bool
	CountWords bool
	InputFile  string
}

func ParseFlags() Options {
	countBytesFlag := flag.Bool("c", false, "Count bytes per file")
	countLinesFlag := flag.Bool("l", false, "Count lines per file")
	countWordsFlag := flag.Bool("w", false, "Count words per file")

	flag.Parse()

	inputFile := ""
	if len(flag.Args()) > 0 {
		inputFile = flag.Args()[0]
	}

	return Options{
		CountBytes: *countBytesFlag,
		CountLines: *countLinesFlag,
		CountWords: *countWordsFlag,
		InputFile:  inputFile,
	}
}
