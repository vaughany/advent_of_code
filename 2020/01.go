// Run with `go run 01.go`

// Build / compile with:
//   env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" 01.go
//   env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" 01.go
//     Using the 'ldflags' option reduces the filesize from 1.7Mb to 1.1Mb with no real change in runtime or compilation speed.
// Pack with:
//   upx --ultra-brute 01
//   upx --ultra-brute 01.exe
//     `upx` will take the filesize down from 1.1Mb to < 400Kb, but will increase the runtime from ~5ms to ~50ms.

package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"strconv"
)

const debug bool = false

var out1, out2 int
var lines []string
var ins []int

// Read a file with many lines and return an array (of strings).
func getInput(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getInstructions(lines []string) []int {
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		ins = append(ins, i)
	}
	return ins
}

func doOutput(o1, o2 int) {
	fmt.Println("Part One: ", o1)
	if o2 != 0 {
		fmt.Println("Part Two: ", o2)
	}
}

func main() {
	msg := "Advent of Code 2020, Day One."
	fmt.Println(msg)
	// log.Println(msg)

	lines = getInput("01.txt")
	ins = getInstructions(lines)

	part1:
	for _, i := range ins {
		for _, j := range ins {
			if i + j == 2020 {
				fmt.Println(i, j, "== 2020!")
				out1 = i * j
				break part1
			}
		}
  }

	part2:
	for _, i := range ins {
		for _, j := range ins {
			for _, k := range ins {
				if i + j + k == 2020 {
					fmt.Println(i, j, k, "== 2020!")
					out2 = i * j * k
					break part2
				}
			}
		}

	}

	// Part One: 1019904
	// Part Two: 176647680
	doOutput(out1, out2)
}
