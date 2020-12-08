// Run with: `go run 06.go`
// Build / compile with: `go build -ldflags "-s -w" 06.go`

// Advent of Code 2020, Day Six.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var debug bool = false
var filename string = "06.txt"
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
	if o1 != 0 {
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

func partOneAndTwo(ins []string) (int, int) {
	groupNum := 0
	groupSize := []int{0}
	groups := []map[rune]int{
		make(map[rune]int),
	}

	for _, i := range ins {
		if debug {
			info(fmt.Sprintf("Instruction: %s", i))
		}
		if i == "" {
			groups = append(groups, make(map[rune]int))
			groupSize = append(groupSize, 0)
			groupNum++
			continue
		}
		for _, ii := range i {
			groups[groupNum][ii]++
		}
		groupSize[groupNum]++
	}

	if debug {
		info(fmt.Sprintf("Group Info: groups: %d, size: %d.", groups, groupSize))
	}

	part1, part2 := 0, 0
	for groupIndex, group := range groups {
		if debug {
			info(fmt.Sprintf("Group Info: group: %d.", group))
		}
		part1 += len(group)
		for _, count := range group {
			if debug {
				info(fmt.Sprintf("Group Count: %d, Group Size: %d.", count, groupSize[groupIndex]))
			}
			if count == groupSize[groupIndex] {
				part2++
			}
		}
	}

	return part1, part2
}

func main() {
	title("Advent of Code 2020, Day Six.")

	var timeSetup, timeOne time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	instructions := getInput(filename)

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 6742
	// Part Two: 3447
	out1, out2 = partOneAndTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
