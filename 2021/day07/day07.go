package aoc2021day7

import (
	"context"
	"math"
)

// Part One: 343468
// Part Two: 96086265
func Part1And2(ctx context.Context, instructions []int) (int, int) {
	var (
		min, max  int
		fuelUsed1 = make(map[int]int)
		fuelUsed2 = make(map[int]int)
	)

	for _, ins := range instructions {
		// 'min' is zero in both cases.
		if ins > max {
			max = ins
		}
	}

	for hPos := min; hPos <= max; hPos++ {
		var (
			fuel1, fuel2 int
		)

		for _, ins := range instructions {
			switch {
			// There's no case for ins == h as nothing needs to be done in that case.
			case ins < hPos:
				fuel1 += hPos - ins
				fuel2 += getFuel(hPos - ins)
			case ins > hPos:
				fuel1 += ins - hPos
				fuel2 += getFuel(ins - hPos)
			}

			// Considerably slower.
			// fuel1 += abs(ins - hPos)
			// fuel2 += getFuel(abs(ins - hPos))
		}
		fuelUsed1[hPos] = fuel1
		fuelUsed2[hPos] = fuel2
	}

	return getLowestFuel(fuelUsed1), getLowestFuel(fuelUsed2)
}

// func abs(in int) int {
// 	if in < 0 {
// 		return -in
// 	}
// 	return in
// }

func getFuel(steps int) int {
	var out int

	for i := steps; i >= 0; i-- {
		out += i
	}

	return out
}

func getLowestFuel(fuelUsed map[int]int) int {
	var out = math.MaxInt

	for _, fuel := range fuelUsed {
		if fuel < out {
			out = fuel
		}
	}

	return out
}
