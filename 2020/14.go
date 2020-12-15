// Run with: `go run 14.go`
// Build / compile with: `go build -ldflags "-s -w" 14.go`

// Advent of Code 2020, Day Fourteen.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var debug bool = false
var filename string = "14.txt"
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
		fmt.Printf("Input file %s had no lines.", filename)
		os.Exit(1)
	}
	if debug {
		info(fmt.Sprintf("Input file %s has %d lines.", filename, len(lines)))
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
	log.Println(string("\u001b[33m") + info + string("\u001b[0m"))
}

func timeinfo(info string) {
	log.Println(string("\u001b[36m") + info + string("\u001b[0m"))
}

func sum(memory [65535]string) int {
	total := 0
	for _, m := range memory {
		d, _ := strconv.ParseInt(m, 2, 64)
		total += int(d)
	}
	return total
}

func partOne(instructions []string) (int) {
	regexAddress := regexp.MustCompile(`\[\d*\]`)
	regexValue := regexp.MustCompile(`\d*$`)

	mask := ""
	address := 0
	value := ""
	memory := [65535]string{"000000000000000000000000000000000000"}

	for _, ins := range instructions {
		if debug {
			fmt.Println(ins)
		}
		if ins[:4] == "mask" {
			mask = ins[7:]
		}
		if ins[:3] == "mem" {
			address, _ = strconv.Atoi(strings.Trim(regexAddress.FindString(ins), "[]"))
			tmp, _ := strconv.Atoi(regexValue.FindString(ins))
			value = fmt.Sprintf("%036b", tmp)

			result := ""
			for index, i := range mask {
				if string(i) == "X" {
					result += string(value[index])
				} else {
					result += string(i)
				}
			}

			memory[address] = result

			if debug {
				fmt.Println("VAL:", value)
				fmt.Println("MSK:", mask)
				fmt.Println("RES:", result, "\n")
			}
		}
	}
	return sum(memory)
}

func partTwo(instructions []string) (int) {
	return -1
}

func main() {
	title("Advent of Code 2020, Day Fourteen.")

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

	// Part One: 13727901897109
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
