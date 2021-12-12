package aoc2021day12

import (
	"context"
	"fmt"
	"regexp"
)

type Octopus struct {
	energy  int
	flashed bool
}

// Part One:
// A simple, recursive DFS solves it easily (says Reddit). https://en.wikipedia.org/wiki/Depth-first_search
func Part1(ctx context.Context, instructions []string) int {
	var (
		output int
	)

	re := regexp.MustCompile(`[a-zA-Z0-9]{1,5}`)
	for _, instruction := range instructions {
		points := re.FindAll([]byte(instruction), 2)
		fmt.Println(string(points[0]), string(points[1]))
	}

	return output
}

// Part Two:
func Part2(ctx context.Context, instructions []string) int {
	var (
		output int
	)

	return output
}
