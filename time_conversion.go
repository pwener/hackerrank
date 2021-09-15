package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	flag := s[len(s)-2:]

	hour := s[0:2]

	intHour, _ := strconv.Atoi(hour)

	if flag == "PM" {
		if intHour != 12 {
			intHour += 12
		}
		return fmt.Sprintf("%s%s", fmt.Sprint(intHour), s[2:len(s)-2])
	}

	if intHour == 12 {
		intHour = 0
	}
	strHour := fmt.Sprint(intHour)
	if len(strHour) == 1 {
		strHour = "0" + strHour
	}
	return fmt.Sprintf("%s%s", strHour, s[2:len(s)-2])
}

func main() {
	fmt.Println(timeConversion("12:40:22AM"))
}
