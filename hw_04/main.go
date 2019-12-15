package main

import (
	"fmt"

	"github.com/KonstantinK3/Go-colearning/hw_04/unimax"
)

func main() {

	data := []int{41, 72, 19, 36, 65}

	result, _ := unimax.FindMax(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	fmt.Printf("Max value: %v\n", result)

	data1 := []string{"cmon", "baby", "light", "my", "fire"}

	result, _ = unimax.FindMax(data1, func(i, j int) bool {
		return data[i] < data[j]
	})

	fmt.Printf("Max value: %v\n", result)

	data2 := []float64{41.1, 72.5, 19.66, 36.0001, 65.2}

	result, _ = unimax.FindMax(data2, func(i, j int) bool {
		return data[i] < data[j]
	})

	fmt.Printf("Max value: %v\n", result)

}
