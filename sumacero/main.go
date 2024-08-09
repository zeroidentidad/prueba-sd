package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumToZero([]int{3, 4, -7, 5, -6, 2, 5, -1, 8}))
}

func sumToZero(arr []int) []int {
	if arr == nil {
		return []int{}
	}

	if len(arr) == 1 && arr[0] == 0 {
		return arr
	}

	for {
		sumMap := make(map[int]int)
		sumMap[0] = -1 // si suma cero está al principio
		sum := 0
		found := false

		for i, num := range arr {
			sum += num
			if idx, exists := sumMap[sum]; exists {
				// quitar subsecuencia que suma cero
				if idx == -1 {
					arr = arr[i+1:]
				} else {
					arr = append(arr[:idx+1], arr[i+1:]...)
				}
				found = true
				break
			}
			sumMap[sum] = i
		}

		if !found {
			break
		}
	}

	// quitar ceros al final
	for len(arr) > 0 && arr[len(arr)-1] == 0 {
		arr = arr[:len(arr)-1]
	}

	return arr
}
