package main
import (
    "os"
    "io"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type Node struct {
    Value int
    Left *Node
    Right *Node
}

type BinaryTree struct {
    Root *Node
}

type Queue struct {
    values []*Node
}

func (q *Queue) Enqueue(v *Node) {
    q.values = append(q.values, v)
}

func (q *Queue) Dequeue() *Node {
    last := q.values[0]
    q.values = q.values[1:]
    return last
}

func (b BinaryTree) OrderTransversal() {
    q := Queue{
        values: []*Node{b.Root},
    }
    
    levelOrder := []int{}

    for len(q.values) != 0 {
        node := q.Dequeue()

        levelOrder = append(levelOrder, node.Value)

        if node.Left != nil { 
            q.Enqueue(node.Left)
        }

        if node.Right != nil {  
            q.Enqueue(node.Right)
        }
    }
    
    for _, v := range levelOrder {
        fmt.Printf("%d ", v)
    }
}

func Add(node *Node, value int) {
    if node.Left == nil && value < node.Value {
        node.Left = &Node{
            Value: value,
        }
        return
    }
    
    if node.Right == nil && value > node.Value {
        node.Right = &Node{
            Value: value,
        }
        return
    }
    
    if value > node.Value {
        Add(node.Right, value)
    } else {
        Add(node.Left, value)
    }
}

func PrintOrderTransversal(i int, nodes []int) {
    b := BinaryTree{
        Root: &Node{
            Value: nodes[0],
        },
    }
    
    for i, n := range nodes {
        if i == 0 {
            continue
        } 
        Add(b.Root, n)
    }
    
    b.OrderTransversal()
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
    
    PrintOrderTransversal(n, lines)
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
