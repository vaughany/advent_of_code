package aoc2021day21

import (
	"context"
	"regexp"
	"strconv"
)

var (
	totalDiceRolls int
)

// Part One: 752247
func Part1(ctx context.Context, instructions []string) int {
	var (
		dice       = rollDeterministicD100()
		p1position int
		p2position int
		p1score    int
		p2score    int
	)

	re := regexp.MustCompile(`[0-9]{1,2}`)
	tmp := re.FindAllString(instructions[0], 2)
	p1position, _ = strconv.Atoi(tmp[1])
	tmp = re.FindAllString(instructions[1], 2)
	p2position, _ = strconv.Atoi(tmp[1])

	for {
		// Player 1
		p1position += dice() % 10
		if p1position > 10 {
			p1position -= 10
		}
		p1score += p1position
		if p1score >= 1000 {
			return totalDiceRolls * p2score
		}

		// Player 2
		p2position += dice() % 10
		if p2position > 10 {
			p2position -= 10
		}
		p2score += p2position
		if p2score >= 1000 {
			return totalDiceRolls * p1score
		}
	}
}

// In theory we're rolling a D100 from 1, incrementing by 1, but the winning score is reached before we have to 'reset' it to 1.
func rollDeterministicD100() func() int {
	dice := 1
	return func() int {
		out := (dice * 3) + 3

		totalDiceRolls += 3
		dice += 3

		return out
	}
}

// Part Two:
func Part2(ctx context.Context, instructions []string) int {
	var (
		output int
	)

	return output
}
