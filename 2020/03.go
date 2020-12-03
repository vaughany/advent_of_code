// Run with: `go run 03.go`
// Build / compile with: `go build -ldflags "-s -w" 03.go`

// Advent of Code 2020, Day Three.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var debug bool = false
var filename string = "03.txt"

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
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

func getInstructions(lines []string) []string {
	tmp := []string{}
	for _, line := range lines {
		tmp = append(tmp, line)
	}
	if debug {
		doGrid(tmp)
	}
	// TODO: check for out-of-range indexes ahead of time and only expand as required. If possible.
	for i := 1; i <= 7; i++ {
		tmp = widen(tmp)
	}
	return tmp
}

func widen(ins []string) []string {
	for i, s := range ins {
		ins[i] = s + s
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

func doGrid(ins []string) {
  for _, i := range ins {
		fmt.Println(i)
	}
	fmt.Println()
}

func traverse(ins []string, xinc int, yinc int) int {
	trees, x, y := 0, 0, 0

	if debug {
		fmt.Printf("===============> Across: %d. Down: %d.\n", xinc, yinc)
	}

	for index, i := range ins {
		if index == y {
			if i[x:x+1] == "#" {
				trees++
				// Find a way to visualise the data without altering it for future runs.
			// 	ins[y] = i[0:x] + "X" + i[x+1:]
			// } else {
			// 	ins[y] = i[0:x] + "O" + i[x+1:]
			}
			if debug {
				fmt.Printf("Movement: %d. Trees hit: %d.\n", index, trees)
			}
			x += xinc
			y += yinc
		}
	}

	return trees
}

func partOne(ins []string) int {
	return traverse(ins, 3, 1)
}

func partTwo(ins []string) int {
	a := traverse(ins, 1, 1)
	b := traverse(ins, 3, 1)
	c := traverse(ins, 5, 1)
	d := traverse(ins, 7, 1)
	e := traverse(ins, 1, 2)
	return a * b * c * d * e
}

func main() {
	title("Advent of Code 2020, Day Three.")

	out1, out2 := 0, 0
	instructions := getInstructions(getInput(filename))

	// Part One: 272
	out1 = partOne(instructions)
	doOutput(out1, out2)

	// Part Two: 3898725600
	out2 = partTwo(instructions)
	doOutput(out1, out2)
}
