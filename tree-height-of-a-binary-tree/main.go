package main

import (
    "os"
    "io"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "math"
)

type Node struct {
    Value int
    Left *Node
    Right *Node
}

type BinaryTree struct {
    Root *Node
    Lenght int
}

func (b *BinaryTree) AddNode(node *Node, value int) {
    if node.Right == nil && node.Value < value {
        node.Right = &Node{Value: value}
        return
    }
    
    if node.Left == nil && node.Value > value {
        node.Left = &Node{Value: value}
        return
    }
    
    if value > node.Value {
        b.AddNode(node.Right, value)
    } else {
        b.AddNode(node.Left, value)
    }
}

func Height(node *Node) int {
    if node.Left == nil && node.Right == nil {
        return 0
    }
    
    left, right := 0, 0
    if node.Left != nil {
        left = Height(node.Left)
    }
    
    if node.Right != nil {
        right = Height(node.Right)
    }
    
    return int(math.Max(float64(left), float64(right))) + 1
}

func CountHeight(n int, nodes []int) {
    b := BinaryTree{
        Root: &Node{Value: nodes[0]},
        Lenght: 0,
    }
    
    for i, n := range nodes {
        // skip root
        if i == 0 {
            continue
        }
        b.AddNode(b.Root, n)
    }
    
    h := Height(b.Root)

    fmt.Println(h)
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int(nTemp)

    var lines = []int{}

    input := strings.Split(readLine(reader), " ")

    for _, i := range input {
        lTemp, err := strconv.ParseInt(i, 10, 64)
        checkError(err)
        lines = append(lines, int(lTemp))
    }
    
    CountHeight(n, lines)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return string(str)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
