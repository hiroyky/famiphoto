package cast

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func Array[T any, V any](list []*T, castFunc func(*T) *V) []*V {
	dst := make([]*V, len(list))
	for i, v := range list {
		dst[i] = castFunc(v)
	}
	return dst
}
