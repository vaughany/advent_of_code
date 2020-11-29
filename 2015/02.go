// Run with: `go run 02.go`
// Build / compile with: `go build -ldflags "-s -w" 02.go`

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var out1, out2 int
var lines []string
var ins []Area

type Area struct {
	X, Y, Z, XY, YZ, ZX, S, P, V, A int
}

// Read a file with many lines and return an array (of strings).
func getInput(filename string) []string {
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
	return lines
}

func getInstructions(lines []string) []Area {
	for _, line := range lines {
		s := strings.Split(line, "x")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		z, _ := strconv.Atoi(s[2])
		ins = append(ins, Area{X: x, Y: y, Z: z})
	}
	return ins
}

func output(o1, o2 int) {
	fmt.Println("Part One:", o1)
	if o2 != 0 {
		fmt.Println("Part Two:", o2)
	}
}

func (s *Area) generateAreaOfBox() {
	s.XY = s.X * s.Y
	s.YZ = s.Y * s.Z
	s.ZX = s.Z * s.X
	s.A = (2 * s.XY) + (2 * s.YZ) + (2 * s.ZX)
}

func (s *Area) generateVolumeOfBox() {
	s.V = s.X * s.Y * s.Z
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

func main() {
	msg := "Advent of Code 2015, Day Two."
	fmt.Println(msg)
	log.Println(msg)

	lines = getInput("02.txt")
	ins = getInstructions(lines)

	for _, i := range ins {
		i.generateAreaOfBox()
		i.discoverSmallestSide()
		out1 += i.A + i.S

		i.generateVolumeOfBox()
		i.discoverSmallestPerimeter()
		out2 += i.V + i.P
	}

	output(out1, out2)
}
