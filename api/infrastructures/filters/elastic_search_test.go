package filters

import (
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPhotoSearchQuery_Body(t *testing.T) {
	t.Run("空白", func(t *testing.T) {
		q := &PhotoSearchQuery{}
		actual := q.Body()
		expected := &es.SearchRequestBody{
			Query: map[string]any{},
			Sort: map[string]any{
				"date_time_original": map[string]any{
					"order": "desc",
				},
			},
		}
		assert.Equal(t, expected, actual)
	})
	t.Run("DateTimeOriginalGTE", func(t *testing.T) {
		q := &PhotoSearchQuery{
			DateTimeOriginalGTE: cast.Ptr(time.Unix(100, 0)),
		}
		actual := q.Body()
		expected := &es.SearchRequestBody{
			Query: map[string]any{
				"range": map[string]any{
					"date_time_original": map[string]any{
						"gte": int64(100000),
					},
				},
			},
			Sort: map[string]any{
				"date_time_original": map[string]any{
					"order": "desc",
				},
			},
		}
		assert.Equal(t, expected, actual)
	})
	t.Run("DateTimeOriginalLT", func(t *testing.T) {
		q := &PhotoSearchQuery{
			DateTimeOriginalLT: cast.Ptr(time.Unix(200, 0)),
		}
		actual := q.Body()
		expected := &es.SearchRequestBody{
			Query: map[string]any{
				"range": map[string]any{
					"date_time_original": map[string]any{
						"lt": int64(200000),
					},
				},
			},
			Sort: map[string]any{
				"date_time_original": map[string]any{
					"order": "desc",
				},
			},
		}
		assert.Equal(t, expected, actual)
	})
	t.Run("DateTimeOriginalGTE, DateTimeOriginalLT", func(t *testing.T) {
		q := &PhotoSearchQuery{
			DateTimeOriginalGTE: cast.Ptr(time.Unix(100, 0)),
			DateTimeOriginalLT:  cast.Ptr(time.Unix(200, 0)),
		}
		actual := q.Body()
		expected := &es.SearchRequestBody{
			Query: map[string]any{
				"range": map[string]any{
					"date_time_original": map[string]any{
						"gte": int64(100000),
						"lt":  int64(200000),
					},
				},
			},
			Sort: map[string]any{
				"date_time_original": map[string]any{
					"order": "desc",
				},
			},
		}
		assert.Equal(t, expected, actual)
	})
}
