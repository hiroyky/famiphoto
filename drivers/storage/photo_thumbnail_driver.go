package storage

import (
	native "errors"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"os"
	"path"
	"path/filepath"
)

func NewPhotoThumbnailDriver() Driver {
	return &photoThumbnailDriver{
		baseDir: path.Join(config.Env.AssetRootPath, "thumbnails"),
	}
}

type photoThumbnailDriver struct {
	baseDir string
}

func (d *photoThumbnailDriver) CreateFile(filePath string, data []byte) error {
	p := path.Join(d.baseDir, filePath)
	dir := filepath.Dir(p)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return f.Close()
}

func (d *photoThumbnailDriver) CreateDir(dirPath string, perm os.FileMode) error {
	panic("Not implmented")
}

func (d *photoThumbnailDriver) Rename(old, file string) error {
	panic("Not implemented")
}

func (d *photoThumbnailDriver) ReadDir(dirPath string) ([]os.FileInfo, error) {
	panic("Not implemented")
}

func (d *photoThumbnailDriver) ReadFile(filePath string) ([]byte, error) {
	panic("Not implemented")
}

func (d *photoThumbnailDriver) Delete(filePath string) error {
	return os.Remove(path.Join(d.baseDir, filePath))
}
func (d *photoThumbnailDriver) DeleteAll(p string) error {
	return os.Remove(path.Join(d.baseDir, p))
}

func (d *photoThumbnailDriver) Glob(pattern string) ([]string, error) {
	panic("Not implemented")
}

func (d *photoThumbnailDriver) Exist(filePath string) bool {
	panic("Not implemented")
}

func (d *photoThumbnailDriver) Stat(filePath string) (os.FileInfo, error) {
	stat, err := os.Stat(path.Join(d.baseDir, filePath))
	if err != nil {
		if native.Is(err, os.ErrNotExist) {
			return nil, errors.New(errors.FileNotFoundError, err)
		}
		return nil, err
	}
	return stat, nil
}
