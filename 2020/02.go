// Run with: `go run 02.go`
// Build / compile with: `go build -ldflags "-s -w" 02.go`

// Advent of Code 2020, Day Two.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var debug bool = false
type Instruction struct {
	min, max int
	subject, password string
}

func init() {
	const (
		defaultDebug = false
		usageDebug = "Display debugging information"
	)
	flag.BoolVar(&debug, "debug", defaultDebug, usageDebug)
	flag.BoolVar(&debug, "d", defaultDebug, usageDebug)
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

func getInstructions(lines []string) []Instruction {
	tmp := []Instruction{}
	for _, line := range lines {
		s1 := strings.Split(line, " ")
		s2 := strings.Split(s1[0], "-")
		min, _ := strconv.Atoi(s2[0])
		max, _ := strconv.Atoi(s2[1])
		tmp = append(tmp, Instruction{min: min, max: max, subject: s1[1][:1], password: s1[2]})
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

func partOne(ins []Instruction) int {
	matches := 0
	for _, i := range ins {
		count := strings.Count(i.password, i.subject)
		if debug {
			info(fmt.Sprintf("%d %d %s %s - %d", i.min, i.max, i.subject, i.password, count))
		}
		if count >= i.min && count <= i.max {
			matches++
		}
	}
	return matches
}

func partTwo(ins []Instruction) int {
	matches := 0
	for _, i := range ins {
		t1 := i.password[i.min-1:i.min]
		t2 := i.password[i.max-1:i.max]
		if debug {
			info(fmt.Sprintf("%d %d %s %s - %s %s", i.min, i.max, i.subject, i.password, t1, t2))
		}
		if (t1 == i.subject && t2 != i.subject) || (t1 != i.subject && t2 == i.subject){
		  matches++
		}
	}
	return matches
}

func main() {
	title("Advent of Code 2020, Day Two.")

	out1, out2 := 0, 0
	instructions := getInstructions(getInput("02.txt"))

	if debug {
		info(fmt.Sprint(instructions))
	}

	// Part One: 572
	out1 = partOne(instructions)
	doOutput(out1, out2)

	// Part Two: 306
	out2 = partTwo(instructions)
	doOutput(out1, out2)
}
