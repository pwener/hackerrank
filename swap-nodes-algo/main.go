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
    Value int32
    Left *Node
    Right *Node
    Level int32
}

type BinaryTree struct {
    Root *Node
}

func NewBinaryTree() *BinaryTree {
    return &BinaryTree{
        Root: &Node{
            Value: 1,
            Level: 1,
        },
    }
}

func (b *BinaryTree) AddLeaves(node *Node, leaves []int32) {
    if leaves[0] != -1 {        
        node.Left = &Node{
            Value: leaves[0],
            Level: node.Level + 1,
        }
    }
    
    if leaves[1] != -1 {
        node.Right = &Node{
            Value: leaves[1],
            Level: node.Level + 1,
        }
    }
}

// last in, first out
type Queue struct {
    values []*Node
}

func (s *Queue) Dequeue() *Node {
    last := s.values[0]
    s.values = s.values[1:]
    return last
}

func (s *Queue) Enqueue(value *Node) {
    s.values = append(s.values, value)
}

func Swap(node *Node, query int32) {
    if node.Level == query || (node.Level % query) == 0 {
        temp := node.Left
        
        node.Left = node.Right
        node.Right = temp
    }
    
    if node.Left != nil {
        Swap(node.Left, query)
    }
    
    if node.Right != nil {
        Swap(node.Right, query)
    }
}

func Inorder(node *Node, arr *[]int32) {
    if node == nil {
        return 
    }

    Inorder(node.Left, arr)
    *arr = append(*arr, node.Value)
    Inorder(node.Right, arr)
}

func (b BinaryTree) Copy() []int32 {
    result := []int32{}

    Inorder(b.Root, &result)
    
    fmt.Printf("result is %v \n", result)

    return result
}


func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
    b := NewBinaryTree()
    q := Queue{values: []*Node{}}

    q.Enqueue(b.Root)

    for _, i := range indexes {
        node := q.Dequeue()

        b.AddLeaves(node, i)

        if node.Left != nil {
            q.Enqueue(node.Left)
        }
        
        if node.Right != nil {
            q.Enqueue(node.Right)
        }
    }

    res := [][]int32{}
    
    for _, query := range queries {
        Swap(b.Root, query)
        r := b.Copy()
        
        res = append(res, r)
    }
    
    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var indexes [][]int32
    for i := 0; i < int(n); i++ {
        indexesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var indexesRow []int32
        for _, indexesRowItem := range indexesRowTemp {
            indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
            checkError(err)
            indexesItem := int32(indexesItemTemp)
            indexesRow = append(indexesRow, indexesItem)
        }

        if len(indexesRow) != 2 {
            panic("Bad input")
        }

        indexes = append(indexes, indexesRow)
    }

    queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var queries []int32

    for i := 0; i < int(queriesCount); i++ {
        queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        queriesItem := int32(queriesItemTemp)
        queries = append(queries, queriesItem)
    }

    result := swapNodes(indexes, queries)

    for i, rowItem := range result {
        for j, colItem := range rowItem {
            fmt.Fprintf(writer, "%d", colItem)

            if j != len(rowItem) - 1 {
                fmt.Fprintf(writer, " ")
            }
        }

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
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
