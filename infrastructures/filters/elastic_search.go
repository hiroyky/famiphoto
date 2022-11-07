package filters

import (
	"bytes"
	"encoding/json"
)

type searchBody map[string]any

func (b searchBody) Buffer() (*bytes.Buffer, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(b); err != nil {
		return nil, err
	}
	return &buf, nil
}

func (b searchBody) MustBuffer() *bytes.Buffer {
	buf, err := b.Buffer()
	if err != nil {
		panic(err)
	}
	return buf
}

type PhotoSearchQuery struct {
	Limit   *int
	Offset  *int
	PhotoID *int
	OwnerID *string
	GroupID *string
	Name    *string
}

func (r *PhotoSearchQuery) Body() searchBody {
	q := map[string]any{}

	if r == nil {
		return q
	}

	if r.Limit != nil {
		q["size"] = *r.Limit
	}
	if r.Offset != nil {
		q["from"] = *r.Offset
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

	q["query"] = map[string]any{
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
