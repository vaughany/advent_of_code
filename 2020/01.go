// Run with `go run 01.go`

// Build / compile with:
//   env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" 01.go
//   env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" 01.go
//     Using the 'ldflags' option reduces the filesize from 1.7Mb to 1.1Mb with no real change in runtime or compilation speed.
// Pack with:
//   upx --ultra-brute 01
//   upx --ultra-brute 01.exe
//     `upx` will take the filesize down from 1.1Mb to < 400Kb, but will increase the runtime from ~5ms to ~50ms.

// Advent of Code 2020, Day One.

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
	"time"
)

var debug bool = false
var filename string = "01.txt"
var timing bool = false

func init() {
	const (
		usageDebug = "Display debugging information"
		usageFilename = "Specify a file to read input from"
	)
	// flag.BoolVar(&debug, "debug", debug, usageDebug)
	flag.BoolVar(&debug, "d", debug, usageDebug)
	// flag.StringVar(&filename, "filename", filename, usageFilename)
	flag.StringVar(&filename, "f", filename, usageFilename)
	flag.BoolVar(&timing, "t", timing, "Display timing information")
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
	}
	if debug {
		info(fmt.Sprintf("Input file has %d lines / instructions.", len(lines)))
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

func timeinfo(info string) {
	log.Println(string("\u001b[36m") + info + string("\u001b[0m"))
}

// Original 'Part One', adding numbers up to 2020.
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

// New 'Part One', based on the 'subtract from 2020' idea.
func partOneB(ins []int) int {
	loops := 0
	for index, i := range ins {
		find := 2020 - i
		for _, j := range ins[index:] {
			loops++
			if j == find {
				if debug {
					info(fmt.Sprintf("%d + %d == 2020! (Loops: %d)", i, j, loops))
				}
				return i * j
			}
		}
	}
	return -1
}

// Original 'Part Two', adding numbers up to 2020.
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

// New 'Part Two', based on the 'subtract from 2020' idea.
func partTwoB(ins []int) int {
	loops := 0
	for iindex, i := range ins {
		find := 2020 - i
		for jindex, j := range ins[iindex:] {
			find2 := find - j
			for _, k := range ins[jindex:] {
				loops++
				if k == find2 {
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

	var timeSetup, timeOne, timeTwo time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	instructions := getInstructions(getInput(filename))

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 1019904
	out1 = partOne(instructions)
	// out1 = partOneB(instructions)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 176647680
	out2 = partTwo(instructions)
	// out2 = partTwoB(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
