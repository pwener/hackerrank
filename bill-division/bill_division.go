package main

import "fmt"

/*
 * Complete the 'bonAppetit' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY bill
 *  2. INTEGER k
 *  3. INTEGER b
 */

func bonAppetit(bill []int32, k int32, b int32) {
	var total int32 = 0

	for i, v := range bill {
		if i != int(k) {
			total += v
		}
	}

	total = total / 2

	if total == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - total)
	}
}

func main() {
	var bill = []int32{3, 10, 2, 9}

	bonAppetit(bill, 1, 12)
}
