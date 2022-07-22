package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/hiroyky/famiphoto/utils/gql"
	"time"
)

const PhotoName = "Photo"

func NewPhotoPagination(result *entities.PhotoSearchResult, limit, offset int) *PhotoPagination {
	pageInfo := newPaginationInfo(result.Total, len(result.Items), limit, offset)
	nodes := cast.Array(result.Items, newPhoto)

	return &PhotoPagination{
		PageInfo: pageInfo,
		Nodes:    nodes,
	}
}

func newPhoto(p *entities.PhotoSearchResultItem) *Photo {
	return &Photo{
		ID:               gql.EncodeIntID(PhotoName, p.PhotoID),
		OwnerID:          gql.EncodeStrID(UserName, p.OwnerID),
		GroupID:          gql.EncodeStrID(GroupName, p.GroupID),
		Name:             p.Name,
		ImportedAt:       p.ImportedAt.Format(time.RFC3339),
		DateTimeOriginal: p.DateTimeOriginal.Format(time.RFC3339),
		PreviewURL:       p.PreviewURL,
		ThumbnailURL:     p.ThumbnailURL,
		FileTypes:        p.FileTypes,
	}
}
