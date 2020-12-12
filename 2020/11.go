// Run with: `go run 11.go`
// Build / compile with: `go build -ldflags "-s -w" 11.go`

// Advent of Code 2020, Day Eleven.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var debug bool = false
var filename string = "11.txt"
var timing bool = false
var occupied string = "#"
var empty string = "L"
var floor string = "."

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

func printGrid(grid []string, run int) {
	fmt.Println("Run:", run)
	for index, line := range grid {
		fmt.Printf("%d\t%s\n", index, line)
	}
	fmt.Println()
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

func getImmediatelyAdjacentSeats(instructions []string, rowi int, coli int, seattype string) int {
	// if debug {
	// 	fmt.Println("row -1:", string(instructions[rowi - 1][coli - 1]), string(instructions[rowi - 1][coli]), string(instructions[rowi - 1][coli + 1]))
	// 	fmt.Println("row:   ", string(instructions[rowi][coli - 1]), "x", string(instructions[rowi][coli+1]))
	// 	fmt.Println("row +1:", string(instructions[rowi + 1][coli - 1]), string(instructions[rowi + 1][coli]), string(instructions[rowi + 1][coli + 1]))
	// 	fmt.Println()
	// }

	width := len(instructions[0]) - 1
	depth := len(instructions) - 1
	matches := 0

	// Row above.
	if rowi >	0 {
		if coli > 0 {
			if string(instructions[rowi - 1][coli - 1]) == seattype {
				matches++
			}
		}
		if string(instructions[rowi - 1][coli]) == seattype {
			matches++
		}
		if coli < width {
			if string(instructions[rowi - 1][coli + 1]) == seattype {
				matches++
			}
		}
	}

	// This row, no middle.
	if coli > 0 {
		if string(instructions[rowi][coli - 1]) == seattype {
			matches++
		}
	}
	if coli < width {
		if string(instructions[rowi][coli + 1]) == seattype {
			matches++
		}
	}

	// Row below.
	if rowi < depth {
		if coli > 0 {
			if string(instructions[rowi + 1][coli - 1]) == seattype {
				matches++
			}
		}
		if string(instructions[rowi + 1][coli]) == seattype {
			matches++
		}
		if coli < width {
			if string(instructions[rowi + 1][coli + 1]) == seattype {
				matches++
			}
		}
	}

	return matches
}

func countOccupiedSeats(instructions []string) int {
	count := 0
	for _, row := range instructions {
		for _, seat := range row {
			if string(seat) == "#" {
				count++
			}
		}
	}
	return count
}

func partOne(instructions []string) (int) {
	runs := 0

	for {
		if debug {
			printGrid(instructions, runs)
		}

		// Create a second array, and store the changes in that.
		newIns := []string{}

		for rowi, row := range instructions {
			newRow := ""
			for coli, seat := range row {

				gs := getImmediatelyAdjacentSeats(instructions, rowi, coli, occupied)

				switch string(seat) {
				case empty:
					if gs == 0 {
						newRow = newRow + occupied
					} else {
						newRow = newRow + empty
					}
				case occupied:
					if gs >= 4 {
						newRow = newRow + empty
					} else {
						newRow = newRow + occupied
					}
				case floor:
					newRow = newRow + floor
				}

			}
			newIns = append(newIns, newRow)
		}

		// Replace instructions array with the newly generated one, and wipe it.
		instructions = newIns
		newIns = []string{}

		// if debug {
		// 	time.Sleep(500 * time.Millisecond)
		// }

		runs++
		// TODO: instead of an arbitrary number of runs: compare this run to the previous and break if identical.
		if runs > 80 {
			if debug {
				fmt.Println("Arbitrary max run limit; breaking.")
			}
			return countOccupiedSeats(instructions)
		}
	}

	return -1
}

func partTwo(instructions []string) (int) {
	return -1
}

func main() {
	title("Advent of Code 2020, Day Eleven.")

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

	// Part One: 2263
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
