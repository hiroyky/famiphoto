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

type PhotoSearchRequest struct {
	Limit   *int64
	Offset  *int64
	PhotoID *int64
}

func (r *PhotoSearchRequest) Body() searchBody {
	q := map[string]any{}

	if r.Limit != nil {
		q["size"] = *r.Limit
	}

	return q
}
