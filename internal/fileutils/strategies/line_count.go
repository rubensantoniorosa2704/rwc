package strategies

import (
	fileutils "github.com/rubensantoniorosa2704/rwc/internal/fileutils/src"
)

type LineCount struct{}

func (w LineCount) Count(line string) int {
	lines, _ := fileutils.CountLines(line)
	return lines
}
