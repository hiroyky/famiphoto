package pagination

func GetOffsetOrDefault(offset *int) int {
	if offset == nil {
		return 0
	}
	return *offset
}

func GetLimitOrDefault(limit *int, defaultLimit, maxLimit int) int {
	if limit == nil {
		return defaultLimit
	}
	if *limit < 0 {
		return 0
	}
	if *limit > maxLimit {
		return maxLimit
	}
	return *limit
}
