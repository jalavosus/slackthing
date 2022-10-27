package utils

import "path/filepath"

func AbsoluteFilePath(filePath string) (string, error) {
	var fp = filePath

	if !filepath.IsAbs(filePath) {
		absPath, err := filepath.Abs(filePath)
		if err != nil {
			return "", err
		}

		fp = absPath
	}

	return fp, nil
}
