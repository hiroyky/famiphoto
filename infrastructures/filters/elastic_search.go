package filters

import (
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/utils/cast"
)

type PhotoSearchQuery struct {
	Limit   *int
	Offset  *int
	PhotoID *int
	OwnerID *string
	GroupID *string
	Name    *string
}

func (r *PhotoSearchQuery) Body() *es.SearchRequestBody {
	q := &es.SearchRequestBody{}

	if r == nil {
		return q
	}

	if r.Limit != nil {
		q.Size = cast.IntPtrToInt64Ptr(r.Limit)
	}
	if r.Offset != nil {
		offset := cast.IntPtrToInt64Ptr(r.Offset)
		q.From = *offset
	}

	mustMatches := make([]map[string]any, 0)
	if r.PhotoID != nil {
		mustMatches = append(mustMatches, map[string]any{"match": map[string]any{"photo_id": *r.PhotoID}})
	}
	if r.OwnerID != nil {
		mustMatches = append(mustMatches, map[string]any{"match": map[string]any{"owner_id": *r.OwnerID}})
	}
	if r.GroupID != nil {
		mustMatches = append(mustMatches, map[string]any{"match": map[string]any{"group_id": *r.GroupID}})
	}

	q.Query = map[string]any{
		"bool": map[string]any{
			"must": mustMatches,
		},
	}

	return q
}

func NewPhotoSearchQuery(id *int, ownerID, groupID *string, limit, offset int) *PhotoSearchQuery {
	q := &PhotoSearchQuery{
		Limit:   &limit,
		Offset:  &offset,
		PhotoID: id,
		OwnerID: ownerID,
		GroupID: groupID,
		Name:    nil,
	}
	return q
}

func NewSinglePhotoSearchQuery(id int) *PhotoSearchQuery {
	limit := 1
	q := &PhotoSearchQuery{
		Limit:   &limit,
		Offset:  nil,
		PhotoID: &id,
		OwnerID: nil,
		GroupID: nil,
		Name:    nil,
	}
	return q
}
