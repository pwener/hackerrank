## Golang running

Sometimes, the problems don't have a main function to read file, as example:

```golang
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
```

## How to run

```bash
cat input01.txt | go run main.go 
```

```bash
cat input01.txt | npx ts-node src/foo.ts
```