package aoc2021day3

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vaughany/advent_of_code/output"
)

// Part One: 3309596
func Part1(ctx context.Context, instructions []string) int {
	var (
		gammaStr   string
		epsilonStr string
	)

	// Loop through each column.
	for j := 0; j < len(instructions[0]); j++ {
		var (
			zeroes int
			ones   int
		)

		// Count the zeroes and ones in each column.
		for _, ins := range instructions {
			if string(ins[j]) == "0" {
				zeroes++
			} else {
				ones++
			}
		}

		// Build a string of binary.
		if zeroes > ones {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}

	// Convert string of binary to decimal.
	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)

	return int(gamma * epsilon)
}

// Part Two: 2981085
func Part2(ctx context.Context, instructions []string) int {
	var (
		oxygen int64
		co2    int64
		debug  = ctx.Value("debug").(bool)
	)

	// Keep an intact copy of the instructions for the CO2 loop, as we're doing destructive things to the instructions.
	instructions2 := instructions

	// The 'oxygen' loop.
	for j := 0; j < len(instructions[0]); j++ {
		var (
			zeroes     int
			ones       int
			zeroesList []int
			onesList   []int
		)

		if debug {
			output.Info(fmt.Sprintf("'Oxygen' instructions: %d", len(instructions)))
		}

		if len(instructions) == 1 {
			break
		}

		for k, ins := range instructions {
			if string(ins[j]) == "0" {
				zeroes++
				zeroesList = append(zeroesList, k)
			} else {
				ones++
				onesList = append(onesList, k)
			}
		}

		switch {
		case zeroes == ones:
			instructions = remove(instructions, zeroesList)
		case zeroes > ones:
			instructions = remove(instructions, onesList)
		case zeroes < ones:
			instructions = remove(instructions, zeroesList)
		}
	}

	// The 'CO2' loop.
	for j := 0; j < len(instructions[0]); j++ {
		var (
			zeroes     int
			ones       int
			zeroesList []int
			onesList   []int
		)

		if debug {
			output.Info(fmt.Sprintf("'CO2' instructions: %d", len(instructions2)))
		}

		if len(instructions2) == 1 {
			break
		}

		for k, ins := range instructions2 {
			if string(ins[j]) == "0" {
				zeroes++
				zeroesList = append(zeroesList, k)
			} else {
				ones++
				onesList = append(onesList, k)
			}
		}

		switch {
		case zeroes == ones:
			instructions2 = remove(instructions2, onesList)
		case zeroes > ones:
			instructions2 = remove(instructions2, zeroesList)
		case zeroes < ones:
			instructions2 = remove(instructions2, onesList)
		}
	}

	oxygen, _ = strconv.ParseInt(instructions[0], 2, 64)
	co2, _ = strconv.ParseInt(instructions2[0], 2, 64)

	return int(oxygen * co2)
}

func remove(instructions []string, removes []int) []string {
	var (
		newInstructions []string
	)

	for k, ins := range instructions {
		found := false
		for _, j := range removes {
			if k == j {
				found = true
				break
			}
		}

		if !found {
			newInstructions = append(newInstructions, ins)
		}
	}

	return newInstructions
}
