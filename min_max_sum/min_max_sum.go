package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	Sort(arr)

	max := arr[0:4]
	min := arr[len(arr)-4 : len(arr)]

	var sumMax int64 = 0
	var sumMin int64 = 0

	for _, i := range max {
		sumMax += int64(i)
	}

	for _, j := range min {
		sumMin += int64(j)
	}

	fmt.Printf(
		"%d %d\n",
		sumMin,
		sumMax,
	)
}

func Sort(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}

	miniMaxSum(arr)
}
