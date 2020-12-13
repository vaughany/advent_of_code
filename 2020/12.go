// Run with: `go run 12.go`
// Build / compile with: `go build -ldflags "-s -w" 12.go`

// Advent of Code 2020, Day Twelve.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

var debug bool = false
var filename string = "12.txt"
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

func turnLeft(dir string, degrees int) string {
	switch dir {
	case "N":
		switch degrees {
		case 90:
			return "W"
		case 180:
			return "S"
		case 270:
			return "E"
		}

	case "E":
		switch degrees {
		case 90:
			return "N"
		case 180:
			return "W"
		case 270:
			return "S"
		}

	case "S":
		switch degrees {
		case 90:
			return "E"
		case 180:
			return "N"
		case 270:
			return "W"
		}

	case "W":
		switch degrees {
		case 90:
			return "S"
		case 180:
			return "E"
		case 270:
			return "N"
		}
	}

	return "x"
}

func turnRight(dir string, degrees int) string {
	switch dir {
	case "N":
		switch degrees {
		case 90:
			return "E"
		case 180:
			return "S"
		case 270:
			return "W"
		}

	case "E":
		switch degrees {
		case 90:
			return "S"
		case 180:
			return "W"
		case 270:
			return "N"
		}

	case "S":
		switch degrees {
		case 90:
			return "W"
		case 180:
			return "N"
		case 270:
			return "E"
		}

	case "W":
		switch degrees {
		case 90:
			return "N"
		case 180:
			return "E"
		case 270:
			return "S"
		}
	}

	return "x"
}

func partOne(instructions []string) (int) {
	x, y, dir := 0, 0, "E"

	for _, instruction := range instructions {
		ins := instruction[:1]
		value, _ := strconv.Atoi(instruction[1:])
		if debug {
			fmt.Printf("%s: %d.\n", ins, value)
		}
		switch ins {
		case "N":
			y += value
		case "S":
			y -= value
		case "E":
			x += value
		case "W":
			x -= value
		case "L":
			dir = turnLeft(dir, value)
		case "R":
			dir = turnRight(dir, value)
		case "F":
			switch dir {
			case "N":
				y += value
			case "S":
				y -= value
			case "E":
				x += value
			case "W":
				x -= value
			}
		}
		if debug {
			fmt.Printf("Dir: %s. X, Y: %d, %d\n\n", dir, x, y)
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func wpleft(value int, wx int, wy int) (int, int) {
	nwx, nwy := 0, 0
	switch value {
	case 90:
		nwx = -wy
		nwy = wx
	case 180:
		nwx = -wx
		nwy = -wy
	case 270:
		nwx = wy
		nwy = -wx
	}
	return nwx, nwy
}

func wpright(value int, wx int, wy int) (int, int) {
	nwx, nwy := 0, 0
	switch value {
	case 90:
		nwx = wy
		nwy = -wx
	case 180:
		nwx = -wx
		nwy = -wy
	case 270:
		nwx = -wy
		nwy = wx
	}
	return nwx, nwy
}

func partTwo(instructions []string) (int) {
	x, y, wx, wy := 0, 0, 10, 1

	for _, instruction := range instructions {
		ins := instruction[:1]
		value, _ := strconv.Atoi(instruction[1:])
		if debug {
			fmt.Printf("%s: %d.\n", ins, value)
		}
		switch ins {
		case "N":
			wy += value
		case "S":
			wy -= value
		case "E":
			wx += value
		case "W":
			wx -= value
		case "L":
			wx, wy = wpleft(value, wx, wy)
		case "R":
			wx, wy = wpright(value, wx, wy)
		case "F":
			x += value * wx
			y += value * wy
		}
		if debug {
			fmt.Printf("X, Y: %d, %d. wX, wY %d, %d.\n\n", x, y, wx, wy)
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func main() {
	title("Advent of Code 2020, Day Twelve.")

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

	// Part One: 1319
	out1 = partOne(instructions)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two: 62434
	out2 = partTwo(instructions)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
