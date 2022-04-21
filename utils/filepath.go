package utils

import (
	"github.com/hiroyky/famiphoto/utils/array"
	"path/filepath"
	"strings"
)

func Directories(path string) []string {
	dir := filepath.Dir(path)
	if dir == "." {
		return []string{}
	}

	return array.Filter(strings.Split(dir, "/"), func(t string) bool {
		return len(t) > 0
	})
}
