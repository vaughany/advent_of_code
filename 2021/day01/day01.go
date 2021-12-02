package aoc2021day1

import "context"

// Part One: 1759
func Part1(ctx context.Context, instructions []int) int {
	return countHigher(ctx, instructions)
}

// Part Two: 1805
func Part2(ctx context.Context, instructions []int) int {
	var (
		numLoops = len(instructions) - 2
		sums     []int
	)

	for i := range instructions {
		if i == numLoops {
			break
		}

		sums = append(sums, instructions[i]+instructions[i+1]+instructions[i+2])
	}

	return countHigher(ctx, sums)
}

func countHigher(ctx context.Context, in []int) int {
	var out int

	for i := range in {
		if i == 0 {
			continue
		}

		if in[i] > in[i-1] {
			out++
		}
	}

	return out
}
