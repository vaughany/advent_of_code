// Run with: `go run 10.go`
// Build / compile with: `go build -ldflags "-s -w" 10.go`

// Advent of Code 2020, Day Ten.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

var debug bool = false
var filename string = "10.txt"
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

func getInstructions(lines []string) []int {
	tmp := []int{}
	for _, line := range lines {
		l, _ := strconv.Atoi(line)
		tmp = append(tmp, l)
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

func partOne(instructions []int) (int) {
	sort.Ints(instructions[:])

	diff1, diff3 := 0, 0

	// Add 0 to the start and 'highest +3' to the end.
	instructions = append([]int{0}, instructions...)
	highest := instructions[len(instructions)-1] + 3
	instructions = append(instructions, []int{highest}...)

	for j := 0; j < len(instructions) - 1; j++ {
		switch instructions[j+1] - instructions[j] {
		case 1:
			diff1++
		case 3:
			diff3++
		}
	}

	if debug {
		info(fmt.Sprintf("1-jolts: %d", diff1))
		info(fmt.Sprintf("3-jolts: %d", diff3))
		info(fmt.Sprintf("%d * %d = %d", diff1, diff3, diff1 * diff3))
	}

	return (diff1 * diff3)
}

func partTwo(instructions []int) (int) {
	return -1
}

func main() {
	title("Advent of Code 2020, Day Ten.")

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

	// Part One: 2664
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two:
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
