package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func beautifulBinaryString(b string) int {
	n := len(b)
	token := "010"
	s := len(token)
	count := 0

	i := 0

	for {
		if i+s > n {
			break
		}

		if b[i:i+s] == token {
			count = count + 1
			i = i + s
		} else {
			i = i + 1
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)

	b := readLine(reader)

	result := beautifulBinaryString(b)

	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
