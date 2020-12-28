// Run with `go run 01.go`
// Build / compile with go build -ldflags "-s -w" 01.go

// Advent of Code 2015, Day One.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

var debug bool = false
var filename string = "01.txt"
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
	fmt.Println("Part One: ", o1)
	fmt.Println("Part Two: ", o2)
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

func partsOneAndTwo(ins []byte) (int, int) {
	floor, firstBasement := 0, 0
	for i, j := range ins {
		switch j {
		case 40:
			floor++
		case 41:
			floor--
		}
		if firstBasement == 0 && floor < 0 {
			firstBasement = i + 1
		}
	}
	return floor, firstBasement
}

func main() {
	title("Advent of Code 2015, Day One.")

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

	// Part One: 138
	// Part Two: 1771
	out1, out2 = partsOneAndTwo(instructions)
	doOutput(out1, out2)

	if timing {
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
