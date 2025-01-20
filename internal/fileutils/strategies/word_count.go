package strategies

import fileutils "github.com/rubensantoniorosa2704/rwc/internal/fileutils/src"

type WordCount struct{}

func (w WordCount) Count(line string) int {
	words, _ := fileutils.CountWords(line)
	return words
}
