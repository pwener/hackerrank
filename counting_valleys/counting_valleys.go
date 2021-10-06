package main

import "fmt"

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	var level, traversed int32 = 0, 0
	var isBellow bool = false

	for _, v := range path {
		if v == 'U' {
			level++
		} else {
			level--
		}

		if level < 0 {
			isBellow = true
		}

		if isBellow && level >= 0 {
			traversed++
			isBellow = false
		}
	}

	return traversed
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))
}
