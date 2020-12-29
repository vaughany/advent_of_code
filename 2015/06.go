// Run with `go run 06.go` or `go fmt 06.go && go run 06.go`
// Build / compile with go build -ldflags "-s -w" 06.go

// Advent of Code 2015, Day Six.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var debug bool = false
var filename string = "06.txt"
var timing bool = false

var gridSize int = 1000

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

func printGrid(grid [][]bool) {
	for i := range grid {
		line := ""
		for j := range grid {
			if grid[i][j] == true {
				line += "@"
			} else {
				line += "."
			}
		}
		info(fmt.Sprint(line))
	}
}

func changeLights(grid [][]bool, x1 int, y1 int, x2 int, y2 int, on bool) [][]bool {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = on
		}
	}
	return grid
}

func toggleLights(grid [][]bool, x1 int, y1 int, x2 int, y2 int) [][]bool {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = !grid[i][j]
		}
	}
	return grid
}

func countLightsOn(grid [][]bool) int {
	out := 0
	for _, i := range grid {
		for _, j := range i {
			if j {
				out++
			}
		}
	}
	return out
}

func partOne(instructions []string) int {
	grid := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}
	if debug {
		printGrid(grid)
	}

	for _, ins := range instructions {
		var x1, y1, x2, y2 int

		switch {
		case strings.HasPrefix(ins, "turn on"):
			fmt.Sscanf(ins, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			changeLights(grid, x1, y1, x2, y2, true)

		case strings.HasPrefix(ins, "turn off"):
			fmt.Sscanf(ins, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			changeLights(grid, x1, y1, x2, y2, false)

		case strings.HasPrefix(ins, "toggle"):
			fmt.Sscanf(ins, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			toggleLights(grid, x1, y1, x2, y2)
		}
	}
	return countLightsOn(grid)
}

func incrementDecrementBrightness(grid [][]int, x1 int, y1 int, x2 int, y2 int, change bool) [][]int {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if change {
				grid[i][j]++
			} else {
				if grid[i][j] > 0 {
					grid[i][j]--
				}
			}
		}
	}
	return grid
}

func doubleIncrementBrightness(grid [][]int, x1 int, y1 int, x2 int, y2 int) [][]int {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] += 2
		}
	}
	return grid
}

func countLightBrightness(grid [][]int) int {
	out := 0
	for _, i := range grid {
		for _, j := range i {
			out += j
		}
	}
	return out
}

func partTwo(instructions []string) int {
	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	for _, ins := range instructions {
		var x1, y1, x2, y2 int

		switch {
		case strings.HasPrefix(ins, "turn on"):
			fmt.Sscanf(ins, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			incrementDecrementBrightness(grid, x1, y1, x2, y2, true)

		case strings.HasPrefix(ins, "turn off"):
			fmt.Sscanf(ins, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			incrementDecrementBrightness(grid, x1, y1, x2, y2, false)

		case strings.HasPrefix(ins, "toggle"):
			fmt.Sscanf(ins, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			doubleIncrementBrightness(grid, x1, y1, x2, y2)
		}
	}
	return countLightBrightness(grid)
}

func main() {
	title("Advent of Code 2015, Day Six.")

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

	// Part One: 543903
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 14687245
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
