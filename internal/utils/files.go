package utils

import (
	"os"
)

func ReadFile(fp string) ([]byte, error) {
	var err error
	
	fp, err = AbsoluteFilePath(fp)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(fp)
}
