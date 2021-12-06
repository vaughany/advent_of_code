package aoc2021day6

import (
	"context"
)

// Part One: 386640
func Part1(ctx context.Context, instructions []int) int {
	return figureOutFishNumbers(instructions, 80)
}

// Part Two: 1733403626279
func Part2(ctx context.Context, instructions []int) int {
	return figureOutFishNumbers(instructions, 256)
}

// This implementation with everything stored in a map appears to be about twice as slow as just using named integers... /shrug
func figureOutFishNumbers2(instructions []int, days int) int {
	var (
		fa = make(map[int]int)
	)

	for _, ins := range instructions {
		fa[ins]++
	}

	for i := 0; i < days; i++ {
		fa[0], fa[1], fa[2], fa[3], fa[4], fa[5], fa[6], fa[7], fa[8] = fa[1], fa[2], fa[3], fa[4], fa[5], fa[6], fa[7]+fa[0], fa[8], fa[0]
	}

	return fa[0] + fa[1] + fa[2] + fa[3] + fa[4] + fa[5] + fa[6] + fa[7] + fa[8]
}

func figureOutFishNumbers(instructions []int, days int) int {
	var (
		zero, one, two, three, four, five, six, seven, eight int
	)

	for _, ins := range instructions {
		switch ins {
		case 1:
			one++
		case 2:
			two++
		case 3:
			three++
		case 4:
			four++
		case 5:
			five++
		}
	}

	for i := 0; i < days; i++ {
		zero, one, two, three, four, five, six, seven, eight = one, two, three, four, five, six, seven+zero, eight, zero
	}

	return zero + one + two + three + four + five + six + seven + eight
}
