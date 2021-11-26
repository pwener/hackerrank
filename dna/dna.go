package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type dna struct {
	start int32
	end   int32
	d     string
}

type vertex struct {
	children    map[rune]int
	leaf        bool
	parent      int
	parentChar  rune
	suffixLink  int
	endWordLink int
	idxs        []int
	pattern     string
	visited 		bool
}

type ahoTrie struct {
	trie []*vertex
	size int
	root int
}

func (aho *ahoTrie) Init() {
	aho.trie = append(aho.trie, &vertex{
		children: make(map[rune]int),
	})
	aho.size++
	aho.root = 0
}

func (aho *ahoTrie) add(pattern string, idx int) {
	currentvertex := aho.root

	for _, v := range pattern {
		if _, ok := aho.trie[currentvertex].children[v]; !ok {
			aho.trie = append(aho.trie, &vertex{
				suffixLink: -1,
				endWordLink: -1,
				parent:     currentvertex,
				parentChar: v,
				idxs:       []int{},
				children:   make(map[rune]int),
				visited: false,
			})

			aho.trie[currentvertex].children[v] = aho.size
			aho.size++
		}

		currentvertex = aho.trie[currentvertex].children[v]
	}

	aho.trie[currentvertex].leaf = true
	aho.trie[currentvertex].pattern = pattern
	aho.trie[currentvertex].idxs = append(aho.trie[currentvertex].idxs, idx)
}

func (aho *ahoTrie) calculateSuffixLink(vertex int) {
	if vertex == aho.root || aho.trie[vertex].parent == aho.root {
		aho.trie[vertex].suffixLink = aho.root
		aho.trie[vertex].endWordLink = aho.root

		if vertex != aho.root && aho.trie[vertex].leaf {
			aho.trie[vertex].endWordLink = vertex
		}

		return
	}

	parentSuffix := aho.trie[aho.trie[vertex].parent].suffixLink

	for {
		suffixNode := aho.trie[parentSuffix]

		if _, ok := suffixNode.children[aho.trie[vertex].parentChar]; ok {
			aho.trie[vertex].suffixLink = suffixNode.children[aho.trie[vertex].parentChar]
			break
		}

		if parentSuffix == aho.root {
			aho.trie[vertex].suffixLink = aho.root
			break
		}

		parentSuffix = suffixNode.suffixLink
	}

	if aho.trie[vertex].leaf {
		aho.trie[vertex].endWordLink = vertex
	} else {
		aho.trie[vertex].endWordLink = aho.trie[aho.trie[vertex].suffixLink].endWordLink
	}
}

type queue struct {
	values []int
	count  int
}

func (q *queue) Init(value int) {
	q.values = make([]int, 0)

	q.values = append(q.values, value)

	q.count = 1
}

// remove old element
func (q *queue) Dequeue() int {
	last := q.values[0]
	q.values = q.values[1:]

	q.count--

	return last
}

func (q *queue) Enqueue(value int) {
	q.values = append(q.values, value)
	q.count++
}

func (aho *ahoTrie) Prepare() {
	vertexQueue := queue{}
	vertexQueue.Init(aho.root)

	for vertexQueue.count > 0 {
		currentvertex := vertexQueue.Dequeue()

		aho.calculateSuffixLink(currentvertex)

		for key := range aho.trie[currentvertex].children {
			curr := aho.trie[currentvertex].children[key]
			if !aho.trie[curr].visited {
				aho.trie[curr].visited = true
				vertexQueue.Enqueue(curr)
			}
		}
	}
}

func (aho *ahoTrie) ProcessString(text string, first int, last int) []int {
	curentState := aho.root

	result := []int{}

	for j := 0; j < len(text); j++ {
		for {
			if _, ok := aho.trie[curentState].children[rune(text[j])]; ok {
				curentState = aho.trie[curentState].children[rune(text[j])]
				break
			}

			if curentState == aho.root {
				break
			}

			curentState = aho.trie[curentState].suffixLink
		}
		checkstate := curentState

		for {
			checkstate = aho.trie[checkstate].endWordLink

			if checkstate == aho.root {
				break
			}

			idxs := aho.trie[checkstate].idxs

			if len(idxs) == 1 {
				idx := idxs[0]

				if idx >= first && idx <= last {
					result = append(result, idx)
				}
			} else {
				for _, idx := range idxs {
					if idx > last {
						break
					}

					if idx >= first {
						result = append(result, idx)
					}
				}
			}

			checkstate = aho.trie[checkstate].suffixLink
		}
	}

	return result
}

func dnaAnalysis(genes []string, health []int32, dnas []*dna) {
	aho := ahoTrie{}
	aho.Init()

	start := time.Now()
	for i, v := range genes {
		aho.add(v, i)
	}

	aho.Prepare()
	elapsed := time.Since(start)
	fmt.Printf("Prepare took %s\n", elapsed)

	min := int(^uint(0) >> 1)
	max := 0

	for _, seq := range dnas {
		matches := aho.ProcessString(seq.d, int(seq.start), int(seq.end))

		total := 0

		for _, idx := range matches {
			total += int(health[idx])
		}

		if total > max {
			max = total
		}

		if total < min {
			min = total
		}
	}

	fmt.Printf("%d %d\n", min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	genesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var genes []string

	for i := 0; i < int(n); i++ {
		genesItem := genesTemp[i]
		genes = append(genes, genesItem)
	}

	healthTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var health []int32

	for i := 0; i < int(n); i++ {
		healthItemTemp, err := strconv.ParseInt(healthTemp[i], 10, 64)
		checkError(err)
		healthItem := int32(healthItemTemp)
		health = append(health, healthItem)
	}

	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	s := int32(sTemp)
	var dnas []*dna

	for sItr := 0; sItr < int(s); sItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		first := int32(firstTemp)

		lastTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		last := int32(lastTemp)

		d := firstMultipleInput[2]
		input := dna{start: first, end: last, d: d}

		dnas = append(dnas, &input)
	}

	dnaAnalysis(
		genes,
		health,
		dnas,
	)
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
