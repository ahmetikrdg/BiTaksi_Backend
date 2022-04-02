package coordinatEvaluation

import (
	"fmt"
)

func KmCalculation(km []float64) float64 {
	var smallest float64 = km[0]
	for _, num := range km {
		if num < smallest {
			smallest = num
		}
	}
	count := len(km)
	fmt.Println("KmCalculation Data:", count)
	return smallest
}
