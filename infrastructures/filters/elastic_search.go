package filters

import (
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/hiroyky/famiphoto/utils/cast"
	"time"
)

type PhotoSearchQuery struct {
	Limit               *int
	Offset              *int
	PhotoID             *int
	Name                *string
	DateTimeOriginalGTE *time.Time
	DateTimeOriginalLT  *time.Time
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

	rangeFilters := make(map[string]any)
	if r.DateTimeOriginalGTE != nil || r.DateTimeOriginalLT != nil {
		dateTimeOriginalRange := make(map[string]any)
		if r.DateTimeOriginalGTE != nil {
			dateTimeOriginalRange["gte"] = (*r.DateTimeOriginalGTE).Unix() * 1000
		}
		if r.DateTimeOriginalLT != nil {
			dateTimeOriginalRange["lt"] = (*r.DateTimeOriginalLT).Unix() * 1000
		}

		rangeFilters["date_time_original"] = dateTimeOriginalRange
	}

	q.Query = map[string]any{}
	if len(mustMatches) > 0 {
		q.Query["bool"] = map[string]any{
			"must": mustMatches,
		}
	}
	if len(rangeFilters) > 0 {
		q.Query["range"] = rangeFilters
	}

	q.Sort = map[string]any{
		"date_time_original": map[string]any{
			"order": "desc",
		},
	}

	return q
}

func NewPhotoSearchQuery(id, year, month, date *int, limit, offset int, exifTimeZone string) *PhotoSearchQuery {
	var dateTimeOriginalGTE *time.Time = nil
	var dateTimeOriginalLT *time.Time = nil
	if year != nil {
		gte := time.Date(*year, time.Month(cast.PtrToVal(month, 1)), cast.PtrToVal(date, 1), 0, 0, 0, 0, utils.MustLoadLocation(exifTimeZone))
		lt := gte.AddDate(1, 0, 0)
		if month != nil {
			lt = gte.AddDate(0, 1, 0)
		}
		if date != nil {
			lt = gte.AddDate(0, 0, 1)
		}
		dateTimeOriginalGTE = cast.Ptr(gte)
		dateTimeOriginalLT = cast.Ptr(lt)
	}

	q := &PhotoSearchQuery{
		Limit:               &limit,
		Offset:              &offset,
		PhotoID:             id,
		Name:                nil,
		DateTimeOriginalGTE: dateTimeOriginalGTE,
		DateTimeOriginalLT:  dateTimeOriginalLT,
	}
	return q
}

func NewSinglePhotoSearchQuery(id int) *PhotoSearchQuery {
	q := &PhotoSearchQuery{
		Limit:   cast.Ptr(1),
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
