package unimax

import (
	"errors"
	"reflect"
)

// Find max value in slice using custom comparator function
func FindMax(slice interface{}, comparator func(i, j int) bool) (interface{}, error) {
	valueType := reflect.ValueOf(slice)
	if valueType.Kind() != reflect.Slice {
		return nil, errors.New("passed not slice variable")
	}

	if valueType.Len() <= 0 {
		return nil, errors.New("passed empty slice")
	}

	maxValueIdx := 0
	for idx := 1; idx < valueType.Len(); idx++ {
		if comparator(maxValueIdx, idx) {
			maxValueIdx = idx
		}
	}

	return valueType.Index(maxValueIdx).Interface(), nil
}
