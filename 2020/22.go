// Run with: `go run 22.go`
// Build / compile with: `go build -ldflags "-s -w" 22.go`

// Advent of Code 2020, Day Twenty Two.

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
var filename string = "22.txt"
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

func getInstructions(instructions []string) ([]int, []int) {
	boolP1, boolP2 := false, false
	player1, player2 := []int{}, []int{}

	for _, ins := range instructions {
		if ins == "" {
			boolP1 = false
			continue
		}
		if boolP1 == true {
			tmp, _ := strconv.Atoi(ins)
			player1 = append(player1, tmp)
		}
		if ins == "Player 1:" {
			boolP1 = true
		}

		if boolP2 == true {
			tmp, _ := strconv.Atoi(ins)
			player2 = append(player2, tmp)
		}
		if ins == "Player 2:" {
			boolP2 = true
		}
	}

	return player1, player2
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

func multiplyAllTheThings(in []int) int {
	// Reverse the slice.
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}

	out := 0
	for index, score := range in {
		out += score * (index + 1)
	}
	return out
}

func partOne(player1 []int, player2 []int) int {
	round := 1
	for {
		if debug {
			info(fmt.Sprintf("-- Round %d --", round))
			info(fmt.Sprintf("Player 1's deck: %v", player1))
			info(fmt.Sprintf("Player 2's deck: %v", player2))
			info(fmt.Sprintf("Player 1 plays: %d", player1[0]))
			info(fmt.Sprintf("Player 2 plays: %d", player2[0]))
		}

		if player1[0] > player2[0] {
			if debug {
				info(fmt.Sprintln("Player 1 wins the round!"))
			}
			// Move the cards to the bottom of the winning pile (in the right order).
			player1 = append(player1, player1[0], player2[0])
		} else {
			if debug {
				info(fmt.Sprintln("Player 2 wins the round!"))
			}
			// Move the cards to the bottom of the winning pile (in the right order).
			player2 = append(player2, player2[0], player1[0])
		}
		// Remove the cards from the top of the piles.
		player1 = player1[1:]
		player2 = player2[1:]

		if len(player1) == 0 {
			if debug {
				info(fmt.Sprint("GAME OVER!!  Player 2 wins!"))
			}
			return multiplyAllTheThings(player2)
		} else if len(player2) == 0 {
			if debug {
				info(fmt.Sprint("GAME OVER!!  Player 1 wins!"))
			}
			return multiplyAllTheThings(player1)
		}

		// if debug {
		// time.Sleep(200 * time.Millisecond)
		// }

		round++
	}

	return -1
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func partTwo(player1 []int, player2 []int) int {
	return -1
}

func main() {
	title("Advent of Code 2020, Day Twenty Two.")

	var timeSetup, timeOne, timeTwo time.Time
	if timing {
		timeSetup = time.Now()
	}

	out1, out2 := 0, 0
	player1, player2 := getInstructions(getInput(filename))

	if timing {
		timeinfo(fmt.Sprintf("Setup took %s", time.Since(timeSetup)))
		timeOne = time.Now()
	}

	// Part One: 31754
	out1 = partOne(player1, player2)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part One took %s", time.Since(timeOne)))
		timeTwo = time.Now()
	}

	// Part Two:
	out2 = partTwo(player1, player2)
	doOutput(out1, out2)
	if timing {
		timeinfo(fmt.Sprintf("Part Two took %s", time.Since(timeTwo)))
		timeinfo(fmt.Sprintf("Both Parts took %s", time.Since(timeOne)))
		timeinfo(fmt.Sprintf("Everything took %s", time.Since(timeSetup)))
	}
}
