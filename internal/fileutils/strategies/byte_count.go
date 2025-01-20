package strategies

import (
	fileutils "github.com/rubensantoniorosa2704/rwc/internal/fileutils/src"
)

type ByteCount struct{}

func (b ByteCount) Count(line string) int {
	bytes, _ := fileutils.CountBytes(line)
	return bytes
}
