package main


import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
type Node struct {
    next *Node
    value int
}

func (n *Node) Add(value int) *Node {
	n.value = value
	n.next = &Node{nil, 0}

	return n
}

func Execute(n int, list []int) {
	head := &Node{
		next: nil,
		value: 0,
	}

	var last *Node = head

	for _, v := range list {
		last.Add(v)
		last = last.next
	}

	it := head

	for it.next != nil {
		fmt.Println(it.value)
		it = it.next
	}
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int(nTemp)

	var lines = []int{}

    for {
		input := strings.TrimSpace(readLine(reader))
		if input == "" {
			break
		}

		lTemp, err := strconv.ParseInt(input, 10, 64)
		checkError(err)

		lines = append(lines, int(lTemp))
	}

	Execute(n, lines)
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