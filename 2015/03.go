// Run with `go run 03.go`
// Build / compile with go build -ldflags "-s -w" 03.go

// Advent of Code 2015, Day Three.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

var debug bool = false
var filename string = "03.txt"
var timing bool = false

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
	flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.Parse()
}

// Read a file with a single line; return bytes.
func getInput(filename string) []byte {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(contents) == 0 {
		panic("Input file has no data.")
	}
	if debug {
		info(fmt.Sprintf("Input file has one line with %d characters.", len(contents)))
	}
	return contents
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

func partOne(ins []byte) int {
	x, y := 0, 0
	grid := map[string]int{}
	for _, j := range ins {
		switch j {
		case 94: // Up.
			y++
		case 118: // Down.
			y--
		case 60: // Left.
			x--
		case 62: // Right.
			x++
		}
		grid[fmt.Sprintf("%d,%d", x, y)]++
		if debug {
			info(fmt.Sprintf("X: %d, Y: %d.", x, y))
		}
	}
	if debug {
		info(fmt.Sprintf("%#v", grid))
	}
	return len(grid)
}

func partTwo(ins []byte) int {
	x1, y1, x2, y2 := 0, 0, 0, 0
	grid := map[string]int{}

	for i, j := range ins {
		if i%2 == 0 {
			// Just do this with runes instead of converting everything to a string and then checking that.
			switch j {
			case 94: // Up.
				y1++
			case 118: // Down.
				y1--
			case 60: // Left.
				x1--
			case 62: // Right.
				x1++
			}
			grid[fmt.Sprintf("%d,%d", x1, y1)]++

		} else {
			switch j {
			case 94: // Up.
				y2++
			case 118: // Down.
				y2--
			case 60: // Left.
				x2--
			case 62: // Right.
				x2++
			}
			grid[fmt.Sprintf("%d,%d", x2, y2)]++
		}
		if debug {
			info(fmt.Sprintf("X1: %d, Y1: %d.", x1, y1))
		}
	}
	if debug {
		info(fmt.Sprintf("%#v", grid))
	}
	return len(grid)
}

func main() {
	title("Advent of Code 2015, Day Three.")

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

	// Part One: 2565
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 2639
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
