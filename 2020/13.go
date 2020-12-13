// Run with: `go run 13.go`
// Build / compile with: `go build -ldflags "-s -w" 13.go`

// Advent of Code 2020, Day Thirteen.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	// "math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var debug bool = false
var filename string = "13.txt"
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

func getInstructions(instructions []string) (int, []string) {
	t, _ := strconv.Atoi(instructions[0])
	b := strings.Split(instructions[1], ",")
	return t, b
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

func partOne(timestamp int, instructions []string) (int) {
	if debug {
		fmt.Println("Ts:", timestamp, "Busses:", instructions)
	}

	// Remove "x".
	newIns := []int{}
	for _, i := range instructions {
		if i != "x" {
			int, _ := strconv.Atoi(i)
			newIns = append(newIns, int)
		}
	}
	sort.Ints(newIns[:])
	if debug {
		fmt.Println("Ts:", timestamp, "Busses:", newIns)
	}

	found := false
	for x := timestamp; x <= timestamp + 10; x++ {
		if debug {
			fmt.Printf("%d:\t", x)
		}

		for _, b := range newIns {
			if x % b == 0 {
				if debug {
					fmt.Printf("%d\t", b)
				}
				found = true
			} else {
				if debug {
					fmt.Print(".\t")
				}
			}
			if found {
				return b * (x - timestamp)
			}
		}
		if debug {
			fmt.Println()
		}
	}

	return -1
}

func partTwo(instructions []string) (int) {
	return -1
}

func main() {
	title("Advent of Code 2020, Day Thirteen.")

	var timeSetup, timeOne, timeTwo time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	instructions := getInput(filename)
	timestamp, busses := getInstructions(instructions)

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 161
	out1 = partOne(timestamp, busses)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 
	out2 = partTwo(busses)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
