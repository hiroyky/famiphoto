package es

import (
	"bytes"
	"encoding/json"
	"io"
)

func (d *driver) Search(index string, body *SearchRequestBody) (*SearchResponse, error) {
	bodyReader, err := body.BodyReader()
	if err != nil {
		return nil, err
	}
	res, err := d.client.Search(
		d.client.Search.WithIndex(index),
		d.client.Search.WithBody(bodyReader),
		d.client.Search.WithTrackScores(true),
	)
	if err := d.handleError(res, err); err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse SearchResponse
	if err := json.Unmarshal(resBody, &searchResponse); err != nil {
		return nil, err
	}
	return &searchResponse, nil
}

type SearchRequestBody struct {
	DocValueFields   any                `json:"docvalue_fields,omitempty"`
	Fields           any                `json:"fields,omitempty"`
	StoredFields     any                `json:"stored_fields,omitempty"`
	Explain          *bool              `json:"explain,omitempty"`
	From             int64              `json:"from"`
	IndicesBoost     map[string]float64 `json:"indices_boost,omitempty"`
	KNN              any                `json:"knn,omitempty"`
	MinScore         *float64           `json:"min_score,omitempty"`
	Pin              any                `json:"pin,omitempty"`
	Query            map[string]any     `json:"query,omitempty"`
	RuntimeMappings  any                `json:"runtime_mappings,omitempty"`
	SeqNoPrimaryTerm *bool              `json:"seq_no_primary_term,omitempty"`
	Size             *int64             `json:"size,omitempty"`
	Sort             map[string]any     `json:"sort,omitempty"`
	Source           any                `json:"_source,omitempty"`
	TerminateAfter   *int64             `json:"terminate_after,omitempty"`
	Timeout          *int64             `json:"timeout,omitempty"`
	Version          *bool              `json:"version,omitempty"`
	Aggs             map[string]any     `json:"aggs,omitempty"`
}

func (r *SearchRequestBody) BodyReader() (io.Reader, error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonBody), nil
}

type SearchResponse struct {
	Took     int64 `json:"took"`
	TimedOut bool  `json:"timed_out"`
	Shards   struct {
		Total      int64 `json:"total"`
		Successful int64 `json:"successful"`
		Skipped    int64 `json:"skipped"`
		Failed     int64 `json:"failed"`
	} `json:"_shards"`
	Hits         Hit `json:"hits"`
	Aggregations map[string]struct {
		Buckets []struct {
			KeyAsString string `json:"key_as_string"`
			Key         int64  `json:"key"`
			DocCount    int64  `json:"doc_count"`
		} `json:"buckets"`
	} `json:"aggregations"`
}

type Hit struct {
	Total struct {
		Value    int64  `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	MaxScore float64    `json:"max_score"`
	Hits     []*HitItem `json:"hits"`
}

type HitItem struct {
	Index  string           `json:"_index"`
	Type   string           `json:"_type"`
	ID     string           `json:"_id"`
	Score  float64          `json:"_score"`
	Source map[string]any   `json:"_source"`
	Fields map[string][]any `json:"fields"`
}
