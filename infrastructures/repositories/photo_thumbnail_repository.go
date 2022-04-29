package repositories

import "github.com/hiroyky/famiphoto/usecases"

func NewPhotoThumbnailRepository(db SQLExecutor) usecases.PhotoThumbnailAdapter {
	return &photoThumbnailRepository{db: db}
}

type photoThumbnailRepository struct {
	db SQLExecutor
}
