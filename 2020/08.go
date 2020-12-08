// Run with: `go run 08.go`
// Build / compile with: `go build -ldflags "-s -w" 08.go`

// Advent of Code 2020, Day Eight.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var debug bool = false
var filename string = "08.txt"
var timing bool = false

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
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
		fmt.Printf("Input file %s had no lines.", filename)
		os.Exit(1)
	}
	if debug {
		info(fmt.Sprintf("Input file %s has %d lines.", filename, len(lines)))
	}
	return lines
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

func isUsedAlready (usedInstructions []int, index int) bool {
	for _, used := range usedInstructions {
		if used == index {
			return true
		}
	}
	return false
}

func partOne(instructions []string) (int) {
	accumulator := 0
	usedInstructions := []int{}

	for index := 0; index < len(instructions); index++ {
		s := strings.Split(instructions[index], " ")
		s2, _ := strconv.Atoi(s[1])

		if debug {
			info(fmt.Sprintf("%d: %s (%d)", index, instructions[index], accumulator))
		}

		if isUsedAlready(usedInstructions, index) {
			return accumulator
		}
		usedInstructions = append(usedInstructions, index)

		switch s[0] {
		case "acc":
			accumulator += s2
		case "jmp":
			index += (s2 -1)
		}

		if debug {
			time.Sleep(10 * time.Millisecond)
		}
	}
	return -1
}

func partTwo(instructions []string) (int) {
	accumulator := 0
  allLoops := 0

	for {
		accumulator = 0
		usedInstructions := []int{}
		jmps := 0
		notThisOne := false

		if debug {
			info(fmt.Sprint("Loop: ", allLoops))
		}

		for index := 0; index < len(instructions); index++ {
			s := strings.Split(instructions[index], " ")
			s2, _ := strconv.Atoi(s[1])

			if debug {
				info(fmt.Sprintf("%d: %s (%d)", index, instructions[index], accumulator))
			}

			if isUsedAlready(usedInstructions, index) {
				notThisOne = true
				break
			}
			usedInstructions = append(usedInstructions, index)

			switch s[0] {
			case "acc":
				accumulator += s2
			case "jmp":
				if jmps == allLoops {
					if debug {
						info(fmt.Sprint("Converted JMP to NOP!"))
					}
				} else {
					index += (s2 - 1)
				}
				jmps++
			}
			if debug {
				time.Sleep(time.Millisecond)
			}

		}
		if notThisOne == false {
			if debug {
				info(fmt.Sprintf("Completed in %d loops: %d.", allLoops, accumulator))
			}
			return accumulator
		}
		allLoops++
	}
	return -1
}

func main() {
	title("Advent of Code 2020, Day Eight.")

	var timeSetup, timeOne, timeTwo time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	instructions := getInput(filename)

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 1137
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 1125
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
