package cast

func BoolToInt[T int](b bool) T {
	if b {
		return 1
	}
	return 0
}
