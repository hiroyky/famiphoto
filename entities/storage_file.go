package entities

import (
	"github.com/hiroyky/famiphoto/utils"
	"github.com/hiroyky/famiphoto/utils/array"
	"strings"
)

type StorageFileInfo struct {
	Name  string
	Path  string
	Ext   string
	IsDir bool
}

func (f StorageFileInfo) IsMatchExt(extensions []string) bool {
	return array.IsContain(strings.ToLower(f.Ext), array.Map(extensions, strings.ToLower))
}

type StorageFileData []byte

func (e StorageFileData) FileHash() string {
	return utils.MD5Bytes(e)
}
