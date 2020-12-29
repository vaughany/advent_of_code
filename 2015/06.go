// Run with `go run 06.go` or `go fmt 06.go && go run 06.go`
// Build / compile with go build -ldflags "-s -w" 06.go

// Advent of Code 2015, Day Six.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var debug bool = false
var filename string = "06.txt"
var timing bool = false

var grid = []string{}
var gridSize int = 1000
var lightOff string = "0" // Grid is too large to represent on a terminal anyway, so these characters hardly matter.
var lightOn string = "1"  // Good to visualise this though.
var repl = strings.NewReplacer(lightOn, lightOff, lightOff, lightOn)

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
	flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.Parse()

	// Grid.
	for x := 0; x <= gridSize-1; x++ {
		grid = append(grid, strings.Repeat(lightOff, 1000))
	}
	if debug {
		printGrid(grid)
	}
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

func getInstructions(instructions []string) []string {
	insRepl := strings.NewReplacer("turn ", "", "through ", "")

	for index, ins := range instructions {
		instructions[index] = insRepl.Replace(ins)
		if debug {
			fmt.Println(instructions[index])
		}
	}
	return instructions
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
	fmt.Println(string("INFO:\t\u001b[33m") + info + string("\u001b[0m"))
}

func timeinfo(info string) {
	fmt.Println(string("DEBUG:\t\u001b[36m") + info + string("\u001b[0m"))
}

func printGrid(in []string) {
	for _, line := range in {
		info(fmt.Sprint(line))
	}
	info("")
}

func changeLights(grid []string, x1 int, y1 int, x2 int, y2 int, on bool) []string {
	char := lightOff
	if on {
		char = lightOn
	}

	for j := x1; j <= x2; j++ {
		insert := strings.Repeat(char, y2-y1+1)
		grid[j] = fmt.Sprint(grid[j][:y1], insert, grid[j][y2+1:])
	}
	return grid
}

func toggleLights(grid []string, x1 int, y1 int, x2 int, y2 int) []string {
	for j := x1; j <= x2; j++ {
		grid[j] = fmt.Sprint(grid[j][:y1], repl.Replace(string(grid[j][y1:y2+1])), grid[j][y2+1:])
	}
	return grid
}

func getCoords(from string, to string) (int, int, int, int) {
	f := strings.Split(from, ",")
	x1, _ := strconv.Atoi(f[0])
	y1, _ := strconv.Atoi(f[1])

	t := strings.Split(to, ",")
	x2, _ := strconv.Atoi(t[0])
	y2, _ := strconv.Atoi(t[1])

	return x1, y1, x2, y2
}

func countLightsOn(grid []string) int {
	out := 0
	for _, line := range grid {
		out += strings.Count(line, lightOn)
	}
	return out
}

func partOne(instructions []string) int {
	for _, ins := range instructions {
		str := strings.Split(ins, " ")
		x1, y1, x2, y2 := getCoords(str[1], str[2])

		switch str[0] {
		case "on":
			changeLights(grid, x1, y1, x2, y2, true)
		case "off":
			changeLights(grid, x1, y1, x2, y2, false)
		case "toggle":
			toggleLights(grid, x1, y1, x2, y2)
		}
	}
	return countLightsOn(grid)
}

func partTwo(ins []string) int {
	return -1
}

func main() {
	title("Advent of Code 2015, Day Six.")

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

	// Part One: 543903
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
