// Run with `go run 04.go`
// Build / compile with go build -ldflags "-s -w" 04.go

// Advent of Code 2015, Day Four.

package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var debug bool = false
var filename string = "04.txt"
var timing bool = false

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
	flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.Parse()
}

// Read a file with a single line; return string.
func getInput(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(bytes) == 0 {
		panic("Input file has no data.")
	}
	bytes = bytes[:len(bytes)-1]
	if debug {
		info(fmt.Sprintf("Input file has one line with %d characters.", len(bytes)))
	}
	return bytes
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
	loop := 0

	for {
		loopBytes := []byte(strconv.Itoa(loop))
		md5 := fmt.Sprintf("%x", md5.Sum(append(ins, loopBytes...)))
		if debug {
			info(fmt.Sprint(md5))
		}

		if strings.HasPrefix(md5, "00000") {
			return loop
		}
		loop++
	}
	return -1
}

func partTwo(ins []byte) int {
	loop := 0

	for {
		loopBytes := []byte(strconv.Itoa(loop))
		md5 := fmt.Sprintf("%x", md5.Sum(append(ins, loopBytes...)))
		if debug {
			info(fmt.Sprint(md5))
		}

		if strings.HasPrefix(md5, "000000") {
			return loop
		}
		loop++
	}
	return -1
}

func main() {
	title("Advent of Code 2015, Day Four.")

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

	// Part One: 282749
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 9962624
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
