package usecases

import (
	"context"
	"fmt"
)

type SearchUseCase interface {
	AppendAllPhotoDocuments(ctx context.Context) error
}

func NewSearchUseCase(
	searchAdapter SearchAdapter,
	photoAdapter PhotoAdapter,
) SearchUseCase {
	return &searchUseCase{
		searchAdapter: searchAdapter,
		photoAdapter:  photoAdapter,
	}
}

type searchUseCase struct {
	searchAdapter SearchAdapter
	photoAdapter  PhotoAdapter
}

func (u *searchUseCase) AppendAllPhotoDocuments(ctx context.Context) error {
	total, err := u.photoAdapter.CountPhotos(ctx)
	if err != nil {
		return err
	}

	limit := int64(500)
	for offset := int64(0); offset <= total; offset += limit {
		photos, err := u.photoAdapter.GetPhotos(ctx, limit, offset)
		if err != nil {
			return err
		}

		files, err := u.photoAdapter.GetPhotoFilesByPhotoIDs(ctx, photos.PhotoIDs())
		if err != nil {
			return err
		}

		stats, err := u.searchAdapter.BulkInsertPhoto(ctx, photos, files, nil)
		if err != nil {
			return err
		}

		fmt.Printf(
			"NumAdded: %d, NumFlushed: %d, NumFailed: %d, NumIndex:%d, NumCreated: %d, NumUpdated:%d, NumDeleted: %d, NumRequest:%d",
			stats.NumAdded,
			stats.NumFlushed,
			stats.NumFailed,
			stats.NumIndexed,
			stats.NumCreated,
			stats.NumUpdated,
			stats.NumDeleted,
			stats.NumDeleted,
		)
	}

	return nil
}
