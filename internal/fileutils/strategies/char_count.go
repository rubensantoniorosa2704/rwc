package strategies

import (
	fileutils "github.com/rubensantoniorosa2704/rwc/internal/fileutils/src"
)

type CharCount struct{}

func (w CharCount) Count(line string) int {
	chars, _ := fileutils.CountChar(line)
	return chars
}
