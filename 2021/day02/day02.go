package aoc2021day2

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Part One: 2120749
func Part1(ctx context.Context, instructions []string) int {
	var (
		depth      int
		horizontal int
	)

	for _, ins := range instructions {

		instruction := strings.Split(ins, " ")
		value, _ := strconv.Atoi(instruction[1])

		switch instruction[0] {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}

		if ctx.Value("debug") == true {
			fmt.Printf("H: %4d.  D: %d.\n", horizontal, depth)
		}
	}

	return horizontal * depth
}

// Part Two: 2138382217
func Part2(ctx context.Context, instructions []string) int {
	var (
		aim        int
		depth      int
		horizontal int
	)

	for _, ins := range instructions {
		instruction := strings.Split(ins, " ")
		value, _ := strconv.Atoi(instruction[1])

		switch instruction[0] {
		case "forward":
			horizontal += value
			depth += (aim * value)
		case "down":
			aim += value
		case "up":
			aim -= value
		}

		if ctx.Value("debug") == true {
			fmt.Printf("H: %4d.  D: %7d.  A:%d.\n", horizontal, depth, aim)
		}
	}

	return horizontal * depth
}
