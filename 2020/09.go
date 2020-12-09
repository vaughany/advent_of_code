// Run with: `go run 09.go`
// Build / compile with: `go build -ldflags "-s -w" 09.go`

// Advent of Code 2020, Day Nine.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var debug bool = false
var filename string = "09.txt"
var timing bool = false
var preamble int = 25
// var preamble int = 5 // Sample data preamble.

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

func doTheseAddUpToThat(in1 int, in2 int, target int) bool {
	if in1 + in2 == target {
		return true
	}
	return false
}

func partOne(instructions []int) (int) {
	for loop := 0; loop <= len(instructions); loop++ {
		found := false
		target := 0

		target = int(instructions[preamble + loop])
		if debug {
			info(fmt.Sprintf("Pre: %d, target: %d.", instructions[loop:preamble + loop], instructions[preamble + loop]))
		}

		ins := instructions[loop:preamble + loop]
FoundAddition:
		for _, j := range ins {
			for _, k := range ins {
				if j == k {
					continue
				}
				if debug {
					fmt.Printf("j: %d + k: %d = %d.\n", j, k, j+k)
				}
				if doTheseAddUpToThat(j, k, target) {
					found = true
					break FoundAddition
				}
			}
		}
		if found == false {
			return target
		}
	}
	return -1
}

func sumThese(array []int) int {
	if len(array) < 2{
		panic("Sum must be minimum of two numbers.")
	}
	res := 0
	for _, a := range array {
		res += a
	}
	return res
}

func smallest(array []int) int {
	res := 0
	for _, a := range array {
		if res == 0 || a < res {
			res = a
		}
	}
	return res
}

func largest(array []int) int {
	res := 0
	for _, a := range array {
		if res == 0 || a > res {
			res = a
		}
	}
	return res
}

func partTwo(instructions []int, target int) (int) {
	if debug {
		info(fmt.Sprint("Target:", target))
	}

	for loop := 0; loop <= len(instructions); loop++ {
		if debug {
			info(fmt.Sprint("Loop:", loop))
		}

		for j := 2; j <= preamble + 1; j++ {
			test := instructions[loop:loop + j]
			answer := sumThese(test)
			if debug {
				fmt.Println(test, answer)
			}
			// Bail early if the sum is greater than the target.
			if answer > target {
				break
			}
			if answer == target {
				return smallest(test) + largest(test)
			}
		}
	}
	return -1
}

func main() {
	title("Advent of Code 2020, Day Nine.")

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

	// Part One: 675280050
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 96081673
	out2 = partTwo(instructions, out1)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
