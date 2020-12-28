// Run with: `go run 02.go`
// Build / compile with: `go build -ldflags "-s -w" 02.go`

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var debug bool = false
var filename string = "02.txt"
var timing bool = false

func init() {
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
	flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.Parse()
}

type Area struct {
	X, Y, Z, XY, YZ, ZX, S, P, V, A int
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

func getInstructions(lines []string) []Area {
	out := []Area{}
	for _, line := range lines {
		s := strings.Split(line, "x")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		z, _ := strconv.Atoi(s[2])
		if debug {
			info(fmt.Sprintf("S: %s, %d %d %d", s, x, y, z))
		}
		out = append(out, Area{X: x, Y: y, Z: z})
	}
	return out
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

func (s *Area) generateAreaOfBox() {
	s.XY = s.X * s.Y
	s.YZ = s.Y * s.Z
	s.ZX = s.Z * s.X
	s.A = (2 * s.XY) + (2 * s.YZ) + (2 * s.ZX)
}

func (s *Area) discoverSmallestSide() {
	var n [3]int
	n[0] = s.XY
	n[1] = s.YZ
	n[2] = s.ZX

	// https://stackoverflow.com/a/53184473/254146
	smallest := n[0]
	for _, num := range n[1:] {
		if num < smallest {
			smallest = num
		}
	}
	s.S = smallest
}

func partOne(ins []Area) int {
	out := 0
	for _, i := range ins {
		i.generateAreaOfBox()
		i.discoverSmallestSide()
		out += i.A + i.S
	}
	return out
}

func (s *Area) generateVolumeOfBox() {
	s.V = s.X * s.Y * s.Z
}

func (s *Area) discoverSmallestPerimeter() {
	var n [3]int
	n[0] = (s.X * 2) + (s.Y * 2)
	n[1] = (s.Y * 2) + (s.Z * 2)
	n[2] = (s.Z * 2) + (s.X * 2)

	smallest := n[0]
	for _, num := range n[1:] {
		if num < smallest {
			smallest = num
		}
	}
	s.P = smallest
}

func partTwo(ins []Area) int {
	out := 0
	for _, i := range ins {
		i.generateVolumeOfBox()
		i.discoverSmallestPerimeter()
		out += i.V + i.P
	}
	return out
}

func main() {
	title("Advent of Code 2015, Day Two.")

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

	// Part One: 1586300
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 3737498
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
