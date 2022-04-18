package array

func IsContain[T comparable](val T, list []T) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}

func Map[T any, V any](list []T, fc func(T) V) []V {
	dst := make([]V, len(list))
	for i, val := range list {
		dst[i] = fc(val)
	}
	return dst
}
