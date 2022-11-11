package usecases

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/golang/mock/gomock"
	"github.com/hiroyky/famiphoto/entities"
	mock_infrastructures "github.com/hiroyky/famiphoto/testing/mocks/infrastructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchUseCase_AppendAllPhotoDocuments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	total := 2100
	photos1 := make(entities.PhotoList, 500)
	photos2 := make(entities.PhotoList, 500)
	photos3 := make(entities.PhotoList, 500)
	photos4 := make(entities.PhotoList, 500)
	photos5 := make(entities.PhotoList, 100)
	status := &esutil.BulkIndexerStats{}

	photoAdapter := mock_infrastructures.NewMockPhotoAdapter(ctrl)
	photoAdapter.EXPECT().CountPhotos(gomock.Any()).Return(total, nil)
	photoAdapter.EXPECT().GetPhotos(gomock.Any(), 500, 0).Return(photos1, nil)
	photoAdapter.EXPECT().GetPhotos(gomock.Any(), 500, 500).Return(photos2, nil)
	photoAdapter.EXPECT().GetPhotos(gomock.Any(), 500, 1000).Return(photos3, nil)
	photoAdapter.EXPECT().GetPhotos(gomock.Any(), 500, 1500).Return(photos4, nil)
	photoAdapter.EXPECT().GetPhotos(gomock.Any(), 500, 2000).Return(photos5, nil)
	searchAdapter := mock_infrastructures.NewMockSearchAdapter(ctrl)
	searchAdapter.EXPECT().BulkInsertPhotos(gomock.Any(), photos1, nil).Return(status, nil)
	searchAdapter.EXPECT().BulkInsertPhotos(gomock.Any(), photos2, nil).Return(status, nil)
	searchAdapter.EXPECT().BulkInsertPhotos(gomock.Any(), photos3, nil).Return(status, nil)
	searchAdapter.EXPECT().BulkInsertPhotos(gomock.Any(), photos4, nil).Return(status, nil)
	searchAdapter.EXPECT().BulkInsertPhotos(gomock.Any(), photos5, nil).Return(status, nil)

	uc := &searchUseCase{
		searchAdapter:  searchAdapter,
		photoAdapter:   photoAdapter,
		appendBulkUnit: 500,
	}
	err := uc.AppendAllPhotoDocuments(context.Background())
	assert.NoError(t, err)
}
