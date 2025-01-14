package fileutils

import (
	"io"
	"os"
	"unicode/utf8"
)

func CountChar(file *os.File) (int, error) {
	buffer := make([]byte, 4096)
	runeCount := 0

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if bytesRead == 0 {
			break
		}

		runeCount += utf8.RuneCount(buffer[:bytesRead])
	}

	return runeCount, nil
}
