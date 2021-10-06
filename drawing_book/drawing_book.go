package main

import (
	"fmt"
)

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n pages long
 *  2. INTEGER p page wanted
 */
func pageCount(n int32, p int32) int32 {
	var final, start int32

	start = lowestPair(p)
	final = lowestPair(n) - lowestPair(p)

	var fromStart float32 = float32(start) / 2
	var fromFinal float32 = float32(final) / 2

	if fromStart < fromFinal {
		return int32(fromStart)
	} else {
		return int32(fromFinal)
	}
}

func lowestPair(n int32) int32 {
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return n
	} else {
		return n - 1
	}
}

func main() {
	fmt.Println(pageCount(6, 5))
}
