package main

import "fmt"

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	s := map[int32]int32{}
	for _, v := range ar {
		s[v] = s[v] + 1
	}

	var sumPairs int32 = 0

	fmt.Println(s)

	for k := range s {
		if s[k]%2 == 0 {
			sumPairs += s[k]
		} else {
			sumPairs += (s[k] - (s[k] % 2))
		}
	}

	return sumPairs / 2
}

func main() {
	var ar []int32 = []int32{1, 2, 1, 2, 1, 3, 2}
	n := int32(7)

	fmt.Println(sockMerchant(n, ar))
}
