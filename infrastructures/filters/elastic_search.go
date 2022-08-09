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

	return q
}

func NewPhotoSearchQuery(id *int, limit, offset int) *PhotoSearchQuery {
	q := &PhotoSearchQuery{
		Limit:   nil,
		Offset:  nil,
		PhotoID: nil,
		OwnerID: nil,
		GroupID: nil,
		Name:    nil,
	}
	return q
}
