package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var EMPTY_RUNE rune = 'w'

type Stack []rune

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

func (s Stack) Push(r rune) Stack {
	return append(s, r)
}

func (s Stack) Pop() (Stack, rune) {
	if s.isEmpty() {
		return s, EMPTY_RUNE
	}
	lastIndex := len(s) - 1
	last := s[lastIndex]

	return s[:lastIndex], last
}

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
func isBalanced(s string) string {
	balancedRel := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	var stack Stack

	for _, c := range s {
		if _, ok := balancedRel[c]; ok {
			fmt.Printf("push %c \n", c)
			stack = stack.Push(c)
		} else {
			var last rune
			stack, last = stack.Pop()

			fmt.Printf("current %c | last %c\n", c, balancedRel[last])

			if last == EMPTY_RUNE {
				return "NO"
			}

			if balancedRel[last] != c {
				return "NO"
			}
		}
	}

	if stack.isEmpty() {
		return "YES"
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Println(result)

		// fmt.Fprintf(writer, "%s\n", result)
	}

	// writer.Flush()
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
