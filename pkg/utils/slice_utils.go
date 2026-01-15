package utils

func FilteredSlice[T any](slice []T, callback func(item T) bool) []T {
	return slice
}

// [1,2,3,4,5]
// FilteredSlice[int]([1,2,3,4,5], func (num int) bool {
//    return num == 3
// })
// [1,2,4,5]

func GetByFieldsSlice[T any](slice []T, callback func(item T) bool) T {
	return slice[0]
}

// [1,2,3,4,5]
// GetByFieldsSlice[int]([1,2,3,4,5], func (num int) bool {
//    return num == 3
// })
// [1,2,4,5]

func UpdateItemSlice[T any](slice []T, newItem T, callback func(item T) bool) bool {
	return false
}

// [1,2,3,4,5]
// UpdateSlice[int]([1,2,3,4,5], 6,func (num int) bool {
//    return num == 5
// })
// [1,2,3,4,5,6]
