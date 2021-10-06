package main

import "fmt"

func main() {
	text := "abcxabsabcxabmnabcxabmnabckm"
	pattern := "abcxabmnabck"

	var i, j int = 0, 0

	lpsArr := mountTableLPS(pattern)

	for i < len(text) && j < len(pattern) {
		if text[i] == pattern[j] {
			i++
			j++
			continue
		}

		if j > 0 {
			j = lpsArr[j - 1]
		} else {
			i++
		}
	}

	if j > len(text) {
		fmt.Println("Not found...")
	} else {
		fmt.Println("Found at index", i - j)
	}
}

/**
* Table of longest possible segments that represents 
* a list of possible fallback positions of pattern to keep searching.
**/
func mountTableLPS(text string) []int {
	relation := make([]int, len(text))
	i, count := 1, 0

	for i < len(text) {
		if text[i] == text[count] {
			count++
			relation[i] = count
			i++
			continue
		} 
		
		if count > 0 {
			count = relation[count-1]
		} else {
			relation[i] = 0
			i++
		}
	}

	return relation
}
