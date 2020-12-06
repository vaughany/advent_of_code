// Run with: `go run 06.go`
// Build / compile with: `go build -ldflags "-s -w" 06.go`

// Advent of Code 2020, Day Six.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	// "strings"
	"time"
)

var debug bool = false
var filename string = "06.txt"
var timing bool = false

type Answer struct {
	// a, b, c, x, y, z, sum int
	sum int
}

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
		info(fmt.Sprintf("Input file has %d lines.", len(lines)))
	}
	return lines
}

func getInstructions(lines []string) []string {
	tmp := []string{}
	for _, line := range lines {
		tmp = append(tmp, line)
	}
	// Add one more blank line.
	tmp = append(tmp, "")
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

func partOne(ins []string) int {
	tmp := []Answer{}
	letters := [26]bool{}
	total := 0

	for _, in := range ins {

		// If there's some data to be processed.
		if in != "" {
			if debug {
				info(fmt.Sprint(in))
			}
			for _, char := range in {
				letters[char - 97] = true
			}

		} else { // Process the data and move on.
			if debug {
				info(fmt.Sprint(">>> end of group"))
			}
			yesses := 0
			for _, value := range letters {
				if value {
					yesses++
				}
			}
			tmp = append(tmp, Answer{sum: yesses})
			letters = [26]bool{}
		}
	}

	for _, i := range tmp {
		total += i.sum
	}
	return total
}

func partTwo(ins []string) int {
	return -1
}

func main() {
	title("Advent of Code 2020, Day Six.")

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

	// Part One: 6742
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
