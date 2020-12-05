// Run with: `go run 05.go`
// Build / compile with: `go build -ldflags "-s -w" 05.go`

// Advent of Code 2020, Day Five.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"time"
)

var debug bool = false
var filename string = "05.txt"
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
		fmt.Println("Input file had no lines.")
		os.Exit(1)
	}
	if debug {
		info(fmt.Sprintf("Input file has %d lines / instructions.", len(lines)))
	}
	return lines
}

func getInstructions(lines []string) []string {
	tmp := []string{}
	for _, line := range lines {
		tmp = append(tmp, line)
	}
	return tmp
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
	mySeatId, highestSeatId, lowestSeatId := 0, 0, 999

	allSeats := []int{}
	for x := 0; x < 1000; x++ {
		allSeats = append(allSeats, x)
	}

	for _, i := range ins {
		j := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(i)
		rowId, _ := strconv.ParseInt(j[:7], 2, 0)
		colId, _ := strconv.ParseInt(j[7:], 2, 0)

		if debug {
			fmt.Println(i, "=", j, ":", rowId, colId)
		}

		seatId := int(rowId) * 8 + int(colId)
		allSeats[seatId] = 0

		if (seatId > highestSeatId) {
			highestSeatId = seatId
		}
		if (seatId < lowestSeatId) {
			lowestSeatId = seatId
		}
	}

	for x := lowestSeatId; x < highestSeatId; x++ {
		if allSeats[x] != 0 {
			mySeatId = allSeats[x]
			break
		}
	}

	return highestSeatId, mySeatId
}

func main() {
	title("Advent of Code 2020, Day Five.")

	var timeSetup, timeOne time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	instructions := getInstructions(getInput(filename))

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 944
	// Part Two: 554
	out1, out2 = partOneAndTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
