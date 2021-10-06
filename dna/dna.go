package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type dna struct {
	start int32
	end   int32
	d     string
}

func dnaAnalysis(genes []string, health []int32, dnas []*dna) {
	// total := []int{}

	seq, gens := findSeq(dnas[1], genes)
	total := calcWeights(seq, gens, health[dnas[1].start:dnas[1].end+1])

	fmt.Println(total)

	// for _, v := range dnas {
	// 	seq, gens := findSeq(v, genes)

	// 	total = append(total, calcWeights(seq, gens, health[v.start:v.end+1]))
	// }

	// sortAsc(total)

	// fmt.Printf("%d %d\n", total[0], total[len(total)-1])
}

func sortAsc(arr []int) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
}

func calcWeights(seq, genes []string, health []int32) int {
	sum := int(0)

	charPointer := make(map[string]int, len(genes))

	for _, v := range genes {
		charPointer[v] = 0
	}

	for _, v := range seq {
		i := charPointer[v] % (len(genes))

		for {
			if v == genes[i] {
				sum += int(health[i])
				charPointer[v] = i + 1
				break
			}
			i = (i + 1) % len(genes)
		}
	}

	return sum
}

func findLimIndexes(genes []string) (int, int) {
	min, max := len(genes[0]), len(genes[0])

	for _, v := range genes {
		if len(v) < min {
			min = len(v)
		}

		if len(v) > max {
			max = len(v)
		}
	}

	return min, max
}

func findSeq(code *dna, genes []string) ([]string, []string) {
	strands := 0

	sequence := []string{}

	genesSlice := genes[code.start : code.end+1]

	lowIndex, highIndex := findLimIndexes(genesSlice)

	fmt.Println("[low, high] is ", lowIndex, highIndex)

	for len(sequence) <= int(code.end-code.start) {
		var target string

		index := lowIndex
		fmt.Println("iteration>>>", strands)

		for index <= highIndex {
			realIndex := strands

			if strands+index > len(code.d) {
				realIndex = len(code.d) - index
			}

			fmt.Printf("checking [%d:%d]\n", realIndex, realIndex+index)

			target = code.d[realIndex : realIndex+index]

			if checkStrInArr(target, genesSlice) {
				sequence = append(sequence, target)
				fmt.Println("appended...")
				break
			}

			index++
		}

		strands++

		fmt.Printf("\n\n")
	}

	fmt.Println(sequence)

	return sequence, genesSlice
}

func checkStrInArr(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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
