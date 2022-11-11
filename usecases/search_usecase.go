package usecases

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
)

type SearchUseCase interface {
	AppendAllPhotoDocuments(ctx context.Context) error
	SearchPhotoByPhotoID(ctx context.Context, id int) (*entities.PhotoSearchResultItem, error)
	SearchPhotos(ctx context.Context, id *int, ownerID, groupID *string, limit, offset int) (*entities.PhotoSearchResult, error)
}

func NewSearchUseCase(
	searchAdapter infrastructures.SearchAdapter,
	photoAdapter infrastructures.PhotoAdapter,
) SearchUseCase {
	return &searchUseCase{
		searchAdapter:  searchAdapter,
		photoAdapter:   photoAdapter,
		appendBulkUnit: 500,
	}
}

type searchUseCase struct {
	searchAdapter  infrastructures.SearchAdapter
	photoAdapter   infrastructures.PhotoAdapter
	appendBulkUnit int
}

func (u *searchUseCase) AppendAllPhotoDocuments(ctx context.Context) error {
	total, err := u.photoAdapter.CountPhotos(ctx)
	if err != nil {
		return err
	}

	limit := u.appendBulkUnit
	for offset := 0; offset <= total; offset += limit {
		photos, err := u.photoAdapter.GetPhotos(ctx, limit, offset)
		if err != nil {
			return err
		}

		stats, err := u.searchAdapter.BulkInsertPhotos(ctx, photos, nil)
		if err != nil {
			return err
		}

		fmt.Printf(
			"NumAdded: %d, NumFlushed: %d, NumFailed: %d, NumIndex:%d, NumCreated: %d, NumUpdated:%d, NumDeleted: %d, NumRequest:%d\n",
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

func (u *searchUseCase) SearchPhotoByPhotoID(ctx context.Context, id int) (*entities.PhotoSearchResultItem, error) {
	query := filters.NewSinglePhotoSearchQuery(id)
	res, err := u.searchAdapter.SearchPhotos(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(res.Items) == 0 {
		return nil, errors.New(errors.DBColumnNotFoundError, nil)
	}
	return res.Items[0], nil
}

func (u *searchUseCase) SearchPhotos(ctx context.Context, id *int, ownerID, groupID *string, limit, offset int) (*entities.PhotoSearchResult, error) {
	query := filters.NewPhotoSearchQuery(id, ownerID, groupID, limit, offset)
	res, err := u.searchAdapter.SearchPhotos(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
