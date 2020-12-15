// Run with: `go run 15.go`
// Build / compile with: `go build -ldflags "-s -w" 15.go`

// Advent of Code 2020, Day Fifteen.

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
var filename string = "15.txt"
var timing bool = false
var output bool = false

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
	flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.BoolVar(&output, "o", output, "Save output to disk")
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

func dumpOutputToFile(filename string, data []int) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w := bufio.NewWriter(file)
	for _, line := range data {
			fmt.Fprintln(w, line)
	}
	// return w.Flush()
	info(fmt.Sprintf("Output file %s saved successully with %d lines.", filename, len(data)))
}

// TODO: One-line input can be handled differently.
func getInsructions(lines []string) []int {
	ins := []int{}
	for _, line := range lines {
		linesArray := strings.Split(line, ",")
		for _, l := range linesArray {
			i, _ := strconv.Atoi(l)
			ins = append(ins, i)
		}
	}
	return ins
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

func partsOneAndTwo(instructions []int, target int) (int) {
	seen := map[int]int{}

	for i, instruction := range instructions {
		seen[instruction] = i + 1
	}
	// As the start of the code deals with the previous number, we kinda forget we ever saw it.
	delete(seen, instructions[len(instructions)-1])

	for {
		key := instructions[len(instructions)-1]
		value := len(instructions)-1

		if seen[key] == 0 {
			if debug {
				fmt.Printf("\nSeen %d for the first time.\n", key)
			}
			seen[key] = value + 1
			instructions = append(instructions, 0)

		} else {
			if debug {
				fmt.Printf("\nSeen %d again.\n", key)
				fmt.Printf("Turn %d - the last time we saw it %d = %d.\n", value + 1, seen[key], value + 1 - seen[key])
			}
			instructions = append(instructions, value + 1 - seen[key])
			seen[key] = value + 1
		}

		if len(instructions) % 10000000 == 0 {
			info(fmt.Sprint("Cycles: ", len(instructions)))
		}

		if len(instructions) == target {
			if output {
				dumpOutputToFile("15-output.txt", instructions)
			}
			return instructions[len(instructions)-1]
		}

		if debug {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(instructions, seen)
		}
	}
	return -1
}

func main() {
	title("Advent of Code 2020, Day Fifteen.")

	var timeSetup, timeOne, timeTwo time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	instructions := getInsructions(getInput(filename))

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 929
	out1 = partsOneAndTwo(instructions, 2020)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 16671510
	out2 = partsOneAndTwo(instructions, 30000000)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
