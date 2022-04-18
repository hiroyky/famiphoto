package entities

import (
	"github.com/hiroyky/famiphoto/utils/array"
	"strings"
)

type StorageFile struct {
	Name  string
	Path  string
	Ext   string
	IsDir bool
}

func (f StorageFile) IsMatchExt(extensions []string) bool {
	return array.IsContain(strings.ToLower(f.Ext), array.Map(extensions, strings.ToLower))
}
