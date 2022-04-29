package utils

import (
	"github.com/hiroyky/famiphoto/utils/array"
	"path/filepath"
	"strings"
)

func SplitPath(path string) []string {
	return array.Filter(strings.Split(path, "/"), func(t string) bool {
		if len(t) == 0 {
			return false
		}
		if t == "." {
			return false
		}
		return true
	})
}

func FileNameExceptExt(p string) string {
	name := filepath.Base(p)
	return strings.TrimSuffix(name, filepath.Ext(name))
}
