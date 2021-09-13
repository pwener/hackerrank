package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	sortDesc(keyboards)
	sortDesc(drives)

	results := []int32{-1}
	for _, dv := range drives {
		for _, kv := range keyboards {
			if kv+dv == b {
				return b
			}

			if kv+dv < b {
				results = append(results, kv+dv)
			}
		}
	}

	sortDesc(results)

	return results[0]
}

func sortDesc(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })
}

func main() {
	keyboards := []int32{
		4,
	}
	drivers := []int32{
		5,
	}
	fmt.Println(getMoneySpent(keyboards, drivers, 5))
}
