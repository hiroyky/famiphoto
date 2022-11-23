package storage

import (
	native "errors"
	"fmt"
	"github.com/hiroyky/famiphoto/errors"
	"os"
	"path"
)

func NewLocalStorageDriver(baseDir string) Driver {
	return &localStorageDriver{baseDir: baseDir}
}

type localStorageDriver struct {
	baseDir string
}

func (d *localStorageDriver) CreateFile(filePath string, data []byte) error {
	panic("Not implemented")
}

func (d *localStorageDriver) CreateDir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(path.Join(d.baseDir, dirPath), perm)
}

func (d *localStorageDriver) Rename(old, file string) error {
	return os.Rename(path.Join(d.baseDir, old), path.Join(d.baseDir, file))
}

func (d *localStorageDriver) ReadDir(dirPath string) ([]os.FileInfo, error) {
	fmt.Println("read dir", path.Join(d.baseDir, dirPath))
	res, err := os.ReadDir(path.Join(d.baseDir, dirPath))
	if err != nil {
		return nil, err
	}
	fileInfoList := make([]os.FileInfo, len(res))
	for i, v := range res {
		fileInfoList[i], err = v.Info()
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, err
		}
	}
	return fileInfoList, nil
}

func (d *localStorageDriver) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(path.Join(d.baseDir, filePath))
}

func (d *localStorageDriver) Delete(filePath string) error {
	return os.Remove(path.Join(d.baseDir, filePath))
}
func (d *localStorageDriver) DeleteAll(p string) error {
	return os.Remove(path.Join(d.baseDir, p))
}

func (d *localStorageDriver) Glob(pattern string) ([]string, error) {
	panic("Not implemented")
}

func (d *localStorageDriver) Exist(filePath string) bool {
	_, err := os.Stat(path.Join(d.baseDir, filePath))
	return err == nil
}

func (d *localStorageDriver) Stat(filePath string) (os.FileInfo, error) {
	stat, err := os.Stat(path.Join(d.baseDir, filePath))
	if err != nil {
		if native.Is(err, os.ErrNotExist) {
			return nil, errors.New(errors.FileNotFoundError, err)
		}
		return nil, err
	}
	return stat, nil
}
