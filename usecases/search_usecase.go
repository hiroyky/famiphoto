package usecases

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/infrastructures"
)

type SearchUseCase interface {
	AppendAllPhotoDocuments(ctx context.Context) error
}

func NewSearchUseCase(
	searchAdapter infrastructures.SearchAdapter,
	photoAdapter infrastructures.PhotoAdapter,
) SearchUseCase {
	return &searchUseCase{
		searchAdapter: searchAdapter,
		photoAdapter:  photoAdapter,
	}
}

type searchUseCase struct {
	searchAdapter infrastructures.SearchAdapter
	photoAdapter  infrastructures.PhotoAdapter
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

		stats, err := u.searchAdapter.BulkInsertPhotos(ctx, photos, nil)
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
