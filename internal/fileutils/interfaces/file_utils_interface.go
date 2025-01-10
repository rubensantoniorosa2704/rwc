package fileutils

import "os"

type FileUtils interface {
	CountBytes(file *os.File) (int, error)
	CountLines(file *os.File) (int, error)
	CountWords(file *os.File) (int, error)
	CountChar(file *os.File) (int, error)
}
