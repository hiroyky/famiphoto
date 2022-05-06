package repositories

import "github.com/hiroyky/famiphoto/utils/array"

func toInterfaceSlice[T any](src []T) []interface{} {
	return array.Map(src, func(t T) interface{} {
		return interface{}(t)
	})
}
