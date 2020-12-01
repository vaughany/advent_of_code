// Run with `go run 01.go`

// Build / compile with:
//   env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" 01.go
//   env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" 01.go
//     Using the 'ldflags' option reduces the filesize from 1.7Mb to 1.1Mb with no real change in runtime or compilation speed.
// Pack with:
//   upx --ultra-brute 01
//   upx --ultra-brute 01.exe
//     `upx` will take the filesize down from 1.1Mb to < 400Kb, but will increase the runtime from ~5ms to ~50ms.

/*
  Version 1.1: added slices of the index of the outer loops to the inner loop, to save repeating the calculations but
	the other way around (part one) or in different orders (part two). Saves about a third of the loops on part one
  (18,710 down from 29,588) and about half of the loops on part two (1,736,373 down from 3,739,732) with my input.
	Also added flags for debugging info, as well as moving the two parts into functions, and general tidying.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var debug bool = false

func init() {
	const (
		defaultDebug = false
		usageDebug = "Display debugging information"
	)
	flag.BoolVar(&debug, "debug", defaultDebug, usageDebug)
	flag.BoolVar(&debug, "d", defaultDebug, usageDebug)
	flag.Parse()
}

// Read a file with many lines and return an array (of strings).
func getInput(filename string) []string {
	lines := []string{}
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
	if len(lines) < 1 {
		fmt.Println("Input file had no lines.")
		os.Exit(1)
	} else {
		if debug {
			info(fmt.Sprintf("Input file has %d lines / instructions.", len(lines)))
		}
	}
	return lines
}

func getInstructions(lines []string) []int {
	tmp := []int{}
	for _, line := range lines {
		t, _ := strconv.Atoi(line)
		tmp = append(tmp, t)
	}
	return tmp
}

func doOutput(o1, o2 int) {
	if o1 != 0 && o2 == 0 {
		fmt.Println("Part One: ", o1)
	}
	if o2 != 0 {
		fmt.Println("Part Two: ", o2)
	}
}

func title(title string) {
	fmt.Println(string("\u001b[32m") + title + string("\u001b[0m"))
}

func info(info string) {
	log.Println(string("\u001b[33m") + info + string("\u001b[0m"))
}

func partOne(ins []int) int {
	loops := 0
	for index, i := range ins {
		for _, j := range ins[index:] {
			loops++
			if i + j == 2020 {
				if debug {
					info(fmt.Sprintf("%d + %d == 2020! (Loops: %d)", i, j, loops))
				}
				return i * j
			}
		}
	}
	return -1
}

func partTwo(ins []int) int {
	loops := 0
	for iindex, i := range ins {
		for jindex, j := range ins[iindex:] {
			for _, k := range ins[jindex:] {
				loops++
				if i + j + k == 2020 {
					if debug {
						info(fmt.Sprintf("%d + %d + %d == 2020! (Loops: %d)", i, j, k, loops))
					}
					return i * j * k
				}
			}
		}
	}
	return -1
}

func main() {
	title("Advent of Code 2020, Day One.")

	out1, out2 := 0, 0
	instructions := getInstructions(getInput("01.txt"))

	// Part One: 1019904
	out1 = partOne(instructions)
	doOutput(out1, out2)

	// Part Two: 176647680
	out2 = partTwo(instructions)
	doOutput(out1, out2)
}
