package model

import "github.com/hiroyky/famiphoto/entities"

func NewDateTimeAggregation(aggregation entities.PhotoDateTimeAggregation) []*DateAggregationItem {
	items := make([]*DateAggregationItem, len(aggregation))
	for i, v := range aggregation {
		items[i] = &DateAggregationItem{
			Year:  v.Year,
			Month: v.Month,
			Date:  v.Date,
			Num:   v.Num,
		}
	}
	return items
}
