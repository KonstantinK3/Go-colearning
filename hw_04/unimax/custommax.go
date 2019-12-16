package unimax

import (
	"errors"
	"reflect"
)

// CustomMax demands slice of anything and coparator function
func CustomMax(input interface{}, comparator func(i, j int) bool) (interface{}, error) {

	inputReflected := reflect.ValueOf(input)
	if inputReflected.Kind() != reflect.Slice {
		return nil, errors.New("Value is not a slice")
	}

	if inputReflected.Len() == 0 {
		return nil, errors.New("Values lenght is not enough")
	}

	maximumIndex := 0
	for id := 0; id < inputReflected.Len(); id++ {
		if comparator(maximumIndex, id) {
			maximumIndex = id
		}
	}

	return inputReflected.Index(maximumIndex).Interface(), nil

}
