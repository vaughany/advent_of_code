// Run with `go run 05.go`
// Build / compile with go build -ldflags "-s -w" 05.go

// Advent of Code 2015, Day Five.

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
var filename string = "05.txt"
var timing bool = false

var banned = []string{"ab", "cd", "pq", "xy"}
var vowels = []string{"a", "e", "i", "o", "u"}
var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var doubles = []string{}
var repeats = []string{}
var pairs = []string{}

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
	flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.Parse()

	// Doubles.
	for _, a1 := range alphabet {
		doubles = append(doubles, fmt.Sprint(a1, a1))
	}
	if debug {
		info(fmt.Sprint(doubles))
	}

	// Repeats.
	for _, a1 := range alphabet {
		for _, a2 := range alphabet {
			repeats = append(repeats, fmt.Sprint(a1, a2, a1))
		}
	}
	if debug {
		info(fmt.Sprint(repeats))
	}

	// Pairs.
	for _, a1 := range alphabet {
		for _, a2 := range alphabet {
			pairs = append(pairs, fmt.Sprint(a1, a2))
		}
	}
	if debug {
		info(fmt.Sprint(pairs))
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

func partOne(ins []string) int {
	out := 0

instructions:
	for index, i := range ins {

		// Filter out the excluded strings first.
		for _, sub := range banned {
			if strings.Contains(i, sub) {
				if debug {
					info(fmt.Sprintf("%d / %s contains '%s'", index, i, sub))
				}
				continue instructions
			}
		}

		// At least three vowels.
		vowelCount := 0
		for _, vowel := range vowels {
			vowelCount += strings.Count(i, vowel)
		}
		if vowelCount < 3 {
			if debug {
				info(fmt.Sprintf("%d / %s contains only %d vowel/s", index, i, vowelCount))
			}
			continue instructions
		}

		// Double letters.
		found := false
		for _, sub := range doubles {
			if strings.Contains(i, sub) {
				found = true
				break
			}
		}
		if found == false {
			if debug {
				info(fmt.Sprintf("%d / %s does not contain any double letters", index, i))
			}
			continue instructions
		}

		// if debug {
		// 	info(fmt.Sprintf("%d / %s is good", index, i))
		// }
		out++
	}

	return out
}

func partTwo(ins []string) int {
	out := 0

instructions:
	for index, i := range ins {

		// "xyx" / "xxx" repeating pattern.
		found := false
		for _, sub := range repeats {
			if strings.Contains(i, sub) {
				found = true
				break
			}
		}
		if found == false {
			if debug {
				info(fmt.Sprintf("%d / %s does not contain any three-letter repeats", index, i))
			}
			continue instructions
		}

		// Pairs of letters, appearing at least twice.
		found2 := false
		for _, pair := range pairs {
			if strings.Count(i, pair) >= 2 {
				found2 = true
				break
			}
		}
		if found2 == false {
			if debug {
				info(fmt.Sprintf("%d / %s does not contain any repeated pairs", index, i))
			}
			continue instructions
		}

		if debug {
			info(fmt.Sprintf("%d / %s is good", index, i))
		}
		out++
	}

	return out
}

func main() {
	title("Advent of Code 2015, Day Five.")

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

	// Part One: 258
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 53
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
