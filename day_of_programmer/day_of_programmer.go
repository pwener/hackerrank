package main

import (
	"fmt"
	"strconv"
)

func divisible(value int32, div int32) bool {
	return value%div == 0
}

func dayOfProgrammer(year int32) string {
	if year == 1918 {
		return "26.09.1918"
	}

	var str string

	if (year < 1917 && divisible(year, 4)) || (year > 1918 && divisible(year, 400) || (divisible(year, 4) && !divisible(year, 100))) {
		str = "12.09." + strconv.Itoa(int(year))
	} else {
		str = "13.09." + strconv.Itoa(int(year))
	}

	return str
}

func main() {
	fmt.Println(dayOfProgrammer(2100))
}
