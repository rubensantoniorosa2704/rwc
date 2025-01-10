package fileutils

import (
	"os"
)

type ByteCount struct {
	File *os.File
}

func (bc *ByteCount) CountBytes() (int, error) {
	fileInfo, err := bc.File.Stat()
	if err != nil {
		return 0, err
	}

	return int(fileInfo.Size()), nil
}
