package main

import (
	"path/filepath"
	"strings"
)

func newPathFormat(localFilePath string) string {
	baseDir = strings.TrimPrefix(baseDir, "./")
	if !strings.HasSuffix(baseDir, "/") {
		baseDir += "/"
	}
	var fileKey string
	if newName != "" {
		fileKey = baseDir + newName
	} else {
		fileKey = baseDir + filepath.Base(localFilePath)
	}
	return fileKey
}
