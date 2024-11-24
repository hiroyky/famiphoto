package entities

type PhotoDateTimeAggregationItem struct {
	Year  int
	Month int
	Date  int
	Num   int
}

type PhotoDateTimeAggregation []*PhotoDateTimeAggregationItem
