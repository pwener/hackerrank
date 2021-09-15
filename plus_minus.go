package main

import "fmt"

func plusMinus(arr []int32) {
	countPlus, countMin, countZero := 0, 0, 0

	for _, v := range arr {
		if v < 0 {
			countMin++
			continue
		}
		if v > 0 {
			countPlus++
			continue
		}

		countZero++
	}

	size := float32(len(arr))

	fmt.Printf("%.6f\n%.6f\n%.6f\n", float32(countPlus)/size, float32(countMin)/size, float32(countZero)/size)
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
