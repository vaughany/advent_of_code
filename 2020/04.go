// Run with: `go run 04.go`
// Build / compile with: `go build -ldflags "-s -w" 04.go`

// Advent of Code 2020, Day Four.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"encoding/hex"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var debug bool = false
var filename string = "04.txt"
var timing bool = false
type Passport struct {
	partOneValidFields, partTwoValidFields int
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
		info(fmt.Sprintf("Input file has %d lines / instructions.", len(lines)))
	}
	return lines
}

func isHexString(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

func isBetween(value string, from int, to int) bool {
	int, _ := strconv.Atoi(value)
	return int >= from && int <= to
}

// Today's input is *potentially* spread over multiple lines (and in no particular order).
func getInstructions(lines []string) []string {
	tmp := []string{}
	tmpstr := ""
	for _, line := range lines {
		if line != "" {
			tmpstr = strings.TrimSpace(fmt.Sprintf("%s %s", tmpstr, line))
		} else {
			tmp = append(tmp, tmpstr)
			tmpstr = ""
		}
	}
	tmp = append(tmp, tmpstr) // Catch the last line.
	return tmp
}

func processInstructions(ins []string) []Passport {
	ppts := []Passport{}
	for _, i := range ins {
		ppt := Passport{}
		fields := strings.Split(i, " ")
		for _, field := range fields {
			fieldTmp := strings.Split(field, ":")
			key := fieldTmp[0]
			value := fieldTmp[1]

			// Part One.
			switch key {
			case "ecl", "hcl", "hgt", "byr", "eyr", "iyr", "pid":
				ppt.partOneValidFields++
			// case "cid":
			// 	ppt.partOneValidFields++
			}

			// Part Two.
			switch key {
			case "ecl":
				switch value {
				case "amb", "blu", "brn", "grn", "gry", "hzl", "oth":
					ppt.partTwoValidFields++
				}
			case "hcl":
				if len(value[1:]) == 6 && isHexString(value[1:]) == true {
					ppt.partTwoValidFields++
				}
			case "hgt":
				measurement := value[len(value)-2:]
				switch measurement {
				case "cm":
					if isBetween(value[:len(value)-2], 150, 193) {
						ppt.partTwoValidFields++
					}
				case "in":
					if isBetween(value[:len(value)-2], 59, 76) {
						ppt.partTwoValidFields++
					}
				}
			case "byr":
				if isBetween(value, 1920, 2002) && len(value) == 4 {
					ppt.partTwoValidFields++
				}
			// case "cid":
			// 	ppt.partTwoValidFields++
			case "eyr":
				if isBetween(value, 2020, 2030) && len(value) == 4 {
					ppt.partTwoValidFields++
				}
			case "iyr":
				if isBetween(value, 2010, 2020) && len(value) == 4 {
					ppt.partTwoValidFields++
				}
			case "pid":
				if len(value) == 9 {
					ppt.partTwoValidFields++
				}
			}
		}
		ppts = append(ppts, ppt)
	}
	return ppts
}

func doOutput(o1, o2 int) {
	if o1 != 0 {
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

func partOneAndTwo(ins []Passport) (int, int) {
	valid1, valid2 := 0, 0
	for _, i := range ins {
		if i.partOneValidFields >= 7 {
			valid1++
		}
		if i.partTwoValidFields >= 7 {
			valid2++
		}
	}
	return valid1, valid2
}

func main() {
	title("Advent of Code 2020, Day Four.")

	var timeSetup, timeOne time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	instructions := processInstructions(getInstructions(getInput(filename)))

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 216
	// Part Two: 150
	out1, out2 = partOneAndTwo(instructions)
	doOutput(out1, out2)
	if timing {
  	timeinfo(fmt.Sprintf("Part One and Two took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
