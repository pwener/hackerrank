// Based on https://www.toptal.com/algorithms/aho-corasick-algorithm
//
// Aho–Corasick algorithm
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
//
// Uses Breadth-first search
// https://en.wikipedia.org/wiki/Breadth-first_search
//
package main

import (
	"fmt"
)

type Vertex struct {
	children map[rune]int
	leaf bool
	parent rune
	parentChar rune
	suffixLink int
	endWordLink int
}

type Aho struct {
	trie []*Vertex
	wordsLen []int
	size int
	root int
}

func (a *Aho) Init() {
	a.trie = append(a.trie, &Vertex{
		children: make(map[rune]int),
	})
	a.size++
	a.root = 0
}

func (aho *Aho) AddString(pattern string) {
	currentVertex := aho.root

	for _, v := range pattern {
		if _, ok := aho.trie[currentVertex].children[v]; !ok {
			aho.trie = append(aho.trie, &Vertex{
				suffixLink: -1,
				parent: rune(currentVertex),
				parentChar: v,
				children: make(map[rune]int),
			})

			aho.trie[currentVertex].children[v] = aho.size
			aho.size++
		}

		currentVertex = aho.trie[currentVertex].children[v];
	}

	aho.trie[currentVertex].leaf = true
	aho.wordsLen = append(aho.wordsLen, len(pattern))
}

func (aho *Aho) calculateSuffLink(vertex int) {
	if (vertex == aho.root) {
		aho.trie[vertex].suffixLink = aho.root
		aho.trie[vertex].endWordLink = aho.root
		return
	}

	if int(aho.trie[vertex].parent) == aho.root {
		aho.trie[vertex].suffixLink = aho.root

		if aho.trie[vertex].leaf {
			aho.trie[vertex].endWordLink = vertex;
		} else {
			aho.trie[vertex].endWordLink = aho.root
			return
		}
	}

	currentBetterVertex := aho.trie[aho.trie[vertex].parent].suffixLink
	chVertex := aho.trie[vertex].parentChar

	for {
		if _, ok := aho.trie[currentBetterVertex].children[chVertex]; ok {
			aho.trie[vertex].suffixLink = aho.trie[currentBetterVertex].children[chVertex]
			break
		}

		if (currentBetterVertex == aho.root) {
			aho.trie[vertex].suffixLink = aho.root;
			break;
		}

		currentBetterVertex = aho.trie[currentBetterVertex].suffixLink;
	}

	if aho.trie[vertex].leaf {
		aho.trie[vertex].endWordLink = vertex;
	} else {
		aho.trie[vertex].endWordLink = aho.trie[aho.trie[vertex].suffixLink].endWordLink
	}
}

type Queue struct {
	values []int
	count int
}

func (q *Queue) Init(value int) {
	q.values = make([]int, 0)

	q.values = append(q.values, value)

	q.count = 1;
}

// remove old element
func (q *Queue) Dequeue() int {
	last := q.values[0]
	q.values = q.values[1:]

	q.count--;

	return last
}

func (q *Queue) Enqueue(value int) {
	q.values = append(q.values, value)
	q.count++;
}


func (aho *Aho) Prepare() {
	vertexQueue := Queue{}
	vertexQueue.Init(aho.root);
	
	for vertexQueue.count > 0 {
		currentVertex := vertexQueue.Dequeue();

		aho.calculateSuffLink(currentVertex)

		for key, _ := range aho.trie[currentVertex].children {
			vertexQueue.Enqueue(aho.trie[currentVertex].children[key])
		}
	}
}

func (aho *Aho) ProcessString(text string) int {
	curentState := aho.root

	result := 0

	for j := 0; j < len(text); j++ {
		for {
			if _, ok := aho.trie[curentState].children[rune(text[j])]; ok {
				curentState = aho.trie[curentState].children[rune(text[j])]
				break;
			}

			if curentState == aho.root {
				break;
			}

			curentState = aho.trie[curentState].suffixLink;
		}
		checkstate := curentState

		for {
			checkstate = aho.trie[checkstate].endWordLink

			if checkstate == aho.root {
				break;
			}

			result++;
			// indexOfMatch = j + 1 - aho.wordsLen[aho.trie[checkstate].workID];

			checkstate = aho.trie[checkstate].suffixLink;
		}
	}

	return result;
}

func main() {
	aho := Aho{}
	aho.Init()

	patterns := []string{"abba", "cab", "baba", "caab", "ac", "abac", "bac" }

	for _, v := range patterns {
		aho.AddString(v);
	}

	aho.Prepare()

	countOfMatches := aho.ProcessString("abbacaabac")

	fmt.Println(countOfMatches)
}