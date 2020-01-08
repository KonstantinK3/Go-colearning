package main

import (
	"fmt"

	"github.com/KonstantinK3/Go-colearning/hw_04/unimax"
)

func main() {

	data := []string{"cmon", "baby", "a", "light", "my", "fire", "z"}
	res := unimax.GetMax(data)
	fmt.Printf("maximum value is: %v \n", res)

	data1 := []float64{1, 2.11, 10, 100, 0, -1}
	res1 := unimax.GetMax(data1)
	fmt.Printf("maximum value is: %v \n", res1)

	data2 := []int{1, 2, 10, 10000, 0, -1}
	res2 := unimax.GetMax(data2)
	fmt.Printf("maximum value is: %v \n", res2)

	data3 := []string{"cmon", "baby", "a", "light", "my", "fire", "z"}
	res3, _ := unimax.CustomMax(data3, func(i, j int) bool {
		return data3[i] < data3[j]
	})
	fmt.Printf("Max value: %v\n", res3)

	data4 := []float64{1, 2.11, 10, 100, 0, -1}
	res4, _ := unimax.CustomMax(data4, func(i, j int) bool {
		return data4[i] < data4[j]
	})
	fmt.Printf("Max value: %v\n", res4)

	data5 := []int{1, 2, 10, 10000, 0, -1}
	res5, _ := unimax.CustomMax(data5, func(i, j int) bool {
		return data5[i] < data5[j]
	})
	fmt.Printf("Max value: %v\n", res5)

}
