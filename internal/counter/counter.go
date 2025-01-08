package counter

import (
	"errors"
)

func CountBytes(content string) (int64, error) {
	if content == "" {
		return 0, errors.New("conteúdo vazio")
	}

	return int64(len(content)), nil
}
