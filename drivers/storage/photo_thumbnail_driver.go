package storage

import (
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"os"
	"path"
	"path/filepath"
)

func NewPhotoThumbnailDriver() repositories.ThumbnailStorageAdapter {
	return &photoThumbnailDriver{
		baseDir: "assets/thumbnails",
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

func (d *photoThumbnailDriver) ReadFile(filePath string) ([]byte, error) {
	panic("Not implemented")
}
func (d *photoThumbnailDriver) Delete(filePath string) error {
	return os.Remove(path.Join(d.baseDir, filePath))
}
func (d *photoThumbnailDriver) DeleteAll(p string) error {
	return os.Remove(path.Join(d.baseDir, p))
}
