package filters

import (
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/hiroyky/famiphoto/utils/cast"
	"time"
)

type PhotoSearchQuery struct {
	Limit   *int
	Offset  *int
	PhotoID *int
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

	q.Query = map[string]any{
		"bool": map[string]any{
			"must": mustMatches,
		},
	}

	q.Sort = map[string]any{
		"date_time_original": map[string]any{
			"order": "desc",
		},
	}

	return q
}

func NewPhotoSearchQuery(id *int, limit, offset int) *PhotoSearchQuery {
	q := &PhotoSearchQuery{
		Limit:   &limit,
		Offset:  &offset,
		PhotoID: id,
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
		Name:    nil,
	}
	return q
}

func NewAggregateByDateTimeOriginalYear(key string) *es.SearchRequestBody {
	return &es.SearchRequestBody{
		Size: cast.Ptr(int64(0)),
		Aggs: map[string]any{
			key: map[string]any{
				"date_histogram": map[string]any{
					"field":             "date_time_original",
					"calendar_interval": "year",
					"format":            "yyyy",
					"min_doc_count":     1,
					"order": map[string]any{
						"_key": "desc",
					},
				},
			},
		},
	}
}

func NewAggregateByDateTimeOriginalYearMonth(key string, year int, tz string) *es.SearchRequestBody {
	locale := utils.MustLoadLocation(tz)
	gte := time.Date(year, 1, 1, 0, 0, 0, 0, locale)
	lt := gte.AddDate(1, 0, 0)

	aggregationYearQuery := &es.SearchRequestBody{
		Query: map[string]any{
			"range": map[string]any{
				"date_time_original": map[string]any{
					"gte": gte.Unix() * 1000,
					"lt":  lt.Unix() * 1000,
				},
			},
		},
		Size: cast.Ptr(int64(0)),
		Aggs: map[string]any{
			key: map[string]any{
				"date_histogram": map[string]any{
					"field":             "date_time_original",
					"calendar_interval": "month",
					"format":            "MM",
					"min_doc_count":     1,
					"order": map[string]any{
						"_key": "asc",
					},
				},
			},
		},
	}

	return aggregationYearQuery
}

func NewAggregateByDateTimeOriginalYearMonthDate(key string, year, month int, tz string) *es.SearchRequestBody {
	locale := utils.MustLoadLocation(tz)
	gte := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, locale)
	lt := gte.AddDate(0, 1, 0)

	aggregationYearQuery := &es.SearchRequestBody{
		Query: map[string]any{
			"range": map[string]any{
				"date_time_original": map[string]any{
					"gte": gte.Unix() * 1000,
					"lt":  lt.Unix() * 1000,
				},
			},
		},
		Size: cast.Ptr(int64(0)),
		Aggs: map[string]any{
			key: map[string]any{
				"date_histogram": map[string]any{
					"field":             "date_time_original",
					"calendar_interval": "day",
					"format":            "dd",
					"min_doc_count":     1,
					"order": map[string]any{
						"_key": "asc",
					},
				},
			},
		},
	}

	return aggregationYearQuery
}
