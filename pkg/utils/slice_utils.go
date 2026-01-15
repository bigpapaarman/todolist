package utils

func FilteredSlice[T any](slice []T, callback func(item T) bool) []T {
	return slice
}

func GetByFieldsSlice[T any](slice []T, callback func(item T) bool) T {
	return slice[0]
}
