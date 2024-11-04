package util

import (
	"fmt"
	"strings"
)

func ConvertUrl(domain, locale, pathName string) string {
	return fmt.Sprintf("https://%v/%v/%v", domain, locale, pathName)
}

func GetFilePathFromFullPath(fullpath string) string {
	splFullPath := strings.Split(fullpath, "/")
	if len(splFullPath) <= 0 {
		return fullpath
	}

	return splFullPath[len(splFullPath)-1]
}