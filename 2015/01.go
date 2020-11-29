// Run with `go run 01.go`

// Build / compile with:
//   env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" 01.go
//   env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" 01.go
//     Using the 'ldflags' option reduces the filesize from 1.7Mb to 1.1Mb with no real change in runtime or compilation speed.
// Pack with:
//   upx --ultra-brute 01
//   upx --ultra-brute 01.exe
//     `upx` will take the filesize down from 1.1Mb to < 400Kb, but will increase the runtime from ~5ms to ~50ms.

// go build -ldflags "-s -w" 01.go && upx --ultra-brute 01 && ls -hl 01 && time ./01

package main

import (
	"bufio"
	"fmt"
	"os"
)

var out1, out2 int
var line string

// Read just one line from a file, 'cos some input is just one line.
func getInput(filename string) string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	line, _ = reader.ReadString('\n')
	return line
}

func output(o1, o2 int) {
	fmt.Println("Part One:", o1)
	if o2 != 0 {
		fmt.Println("Part Two:", o2)
	}
}

func main() {
	fmt.Println("Advent of Code 2015, Day One.")

	line = getInput("01.txt")

	for i, character := range line {
		if string(character) == "(" {
			out1++
		} else {
			out1--
		}

		if out2 == 0 && out1 == -1 {
			out2 = i + 1
		}
	}

	output(out1, out2)
}
