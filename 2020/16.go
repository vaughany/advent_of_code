// Run with: `go run 16.go`
// Build / compile with: `go build -ldflags "-s -w" 16.go`

// Advent of Code 2020, Day Sixteen.

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
var filename string = "16.txt"
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

func getValid(valid map[int]bool, input string) map[int]bool {
	ranges := strings.Split(input, " or ")
	for i, rng := range ranges {
		if debug {
			info(fmt.Sprintf("Range %d: %#v", i, rng))
		}
		r := strings.Split(rng, "-")
		r0, _ := strconv.Atoi(r[0])
		r1, _ := strconv.Atoi(r[1])
		if debug {
			info(fmt.Sprintf("Range: %#v to %#v", r0, r1))
		}
		for x := r0; x <= r1; x++ {
			valid[x] = true
		}
	}

	return valid
}

func sum(in []int) int {
	out := 0
	for _, i := range in {
		out += i
	}
	return out
}

func partOne(instructions []string) int {
	regexClass := regexp.MustCompile(`^class: [0-9]{1,3}-[0-9]{1,3}.*[0-9]{1,3}-[0-9]{1,3}`)
	regexRow := regexp.MustCompile(`^row: [0-9]{1,3}-[0-9]{1,3}.*[0-9]{1,3}-[0-9]{1,3}`)
	regexSeat := regexp.MustCompile(`^seat: [0-9]{1,3}-[0-9]{1,3}.*[0-9]{1,3}-[0-9]{1,3}`)

	class, row, seat, yourTicket := "", "", "", ""
	nearbyTickets := []string{}
	boolYourTicket := false
	boolNearbyTickets := false

	for _, ins := range instructions {
		if class == "" {
			class = regexClass.FindString(ins)
		}
		if row == "" {
			row = regexRow.FindString(ins)
		}
		if seat == "" {
			seat = regexSeat.FindString(ins)
		}

		if boolYourTicket == true {
			yourTicket = ins
			boolYourTicket = false
		}
		if ins == "your ticket:" {
			boolYourTicket = true
		}

		if boolNearbyTickets == true {
			nearbyTickets = append(nearbyTickets, ins)
		}
		if ins == "nearby tickets:" {
			boolNearbyTickets = true
		}
	}

	class = class[7:]
	row = row[5:]
	seat = seat[6:]

	if debug {
		info(fmt.Sprintf("Class: %#v", class))
		info(fmt.Sprintf("Row: %#v", row))
		info(fmt.Sprintf("Seat: %#v", seat))
		info(fmt.Sprintf("Your Ticket: %#v", yourTicket))
		info(fmt.Sprintf("Nearby Tickets: %#v", nearbyTickets))
	}

	valid := map[int]bool{}
	valid = getValid(valid, class)
	valid = getValid(valid, row)
	valid = getValid(valid, seat)

	if debug {
		info(fmt.Sprintf("Valid: %#v", valid))
	}

	invalid := []int{}
	for _, nTs := range nearbyTickets {
		nnt := []int{}

		nt := strings.Split(nTs, ",")
		for _, n := range nt {
			nx, _ := strconv.Atoi(n)
			nnt = append(nnt, nx)
		}

		for _, n := range nnt {
			if valid[n] == false {
				invalid = append(invalid, n)
			}
		}
	}

	return sum(invalid)
}

func partTwo(instructions []string) int {
	return -1
}

func main() {
	title("Advent of Code 2020, Day Sixteen.")

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

	// Part One: 27870
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
