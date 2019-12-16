package unimax

import (
	"reflect"
)

func GetMax(input interface{}) interface{} {

	inputReflected := reflect.ValueOf(input)

	sliceString, ok := inputReflected.Interface().([]string)
	if ok == true {
		max := ""
		for _, value := range sliceString {
			if value > max {
				max = value
			}
		}
		return max
	}

	sliceInt, ok := inputReflected.Interface().([]int)
	if ok == true {
		max := 0
		for _, value := range sliceInt {
			if value > max {
				max = value
			}
		}
		return max
	}

	sliceFloat, ok := inputReflected.Interface().([]float64)
	if ok == true {
		max := 0.0
		for _, value := range sliceFloat {
			if value > max {
				max = value
			}
		}
		return max
	}
	panic("unsupported type")

	// return nil, nil

}
