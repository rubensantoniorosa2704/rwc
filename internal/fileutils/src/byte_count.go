package fileutils

import (
	"os"
)

func CountBytes(file *os.File) (int, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return int(fileInfo.Size()), nil
}
