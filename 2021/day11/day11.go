package aoc2021day11

import (
	"context"
	"strconv"
	"strings"
)

type Octopus struct {
	energy  int
	flashed bool
}

// Part One: 1675
func Part1(ctx context.Context, instructions []string) int {
	var (
		flashes, newFlashes int
	)

	grid := makeGrid(instructions)

	for s := 1; s <= 100; s++ {
		grid, newFlashes = countFlashesAndResetEnergy(checkEnergy(evolveGrid(grid)))
		flashes += newFlashes
	}

	return flashes
}

// Part Two: 515
func Part2(ctx context.Context, instructions []string) int {
	var (
		newFlashes int
		count      int = 1
	)

	grid := makeGrid(instructions)

	for {
		grid, newFlashes = countFlashesAndResetEnergy(checkEnergy(evolveGrid(grid)))
		if newFlashes == 100 {
			return count
		}

		count++
	}
}

func makeGrid(instructions []string) [][]Octopus {
	var grid [][]Octopus

	for _, ins := range instructions {
		tmp := []Octopus{}
		str := strings.Split(ins, "")
		for _, s := range str {
			i, _ := strconv.Atoi(s)
			tmp = append(tmp, Octopus{i, false})
		}
		grid = append(grid, tmp)
	}

	return grid
}

func evolveGrid(grid [][]Octopus) [][]Octopus {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x].energy++
		}
	}

	return grid
}

func checkEnergy(grid [][]Octopus) [][]Octopus {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].energy > 9 {
				grid[y][x].energy = 0
				grid[y][x].flashed = true
				grid = updateAdjacent(grid, x, y)
			}
		}
	}

	return grid
}

func updateAdjacent(grid [][]Octopus, xx, yy int) [][]Octopus {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			// Don't update the originator again.
			if x == 0 && y == 0 {
				continue
			}

			// Bounds check.
			if xx+x < 0 || xx+x > 9 || yy+y < 0 || yy+y > 9 {
				continue
			}

			// if !grid[yy+y][xx+x].flashed {
			grid[yy+y][xx+x].energy++
			// }

			if grid[yy+y][xx+x].energy > 9 {
				grid[yy+y][xx+x].energy = 0
				grid[yy+y][xx+x].flashed = true
				grid = updateAdjacent(grid, xx+x, yy+y)
			}
		}
	}

	return grid
}

func countFlashesAndResetEnergy(grid [][]Octopus) ([][]Octopus, int) {
	var flashes int

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].flashed {
				grid[y][x].flashed = false
				grid[y][x].energy = 0
				flashes++
			}
		}
	}

	return grid, flashes
}
