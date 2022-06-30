package storage

import "os"

type Driver interface {
	CreateFile(filePath string, data []byte) error
	CreateDir(dirPath string, perm os.FileMode) error
	Rename(old, file string) error
	ReadDir(dirPath string) ([]os.FileInfo, error)
	ReadFile(filePath string) ([]byte, error)
	Glob(pattern string) ([]string, error)
	Exist(filePath string) bool
	Delete(filePath string) error
	DeleteAll(path string) error
}
