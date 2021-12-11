package aoc2021day9

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	height int
	lowest bool
}

// Part One:
func Part1(ctx context.Context, instructions []string) int {
	var (
		output int
		grid   [][]Point
	)

	// Create the grid.
	for _, ins := range instructions {
		tmp := []Point{}
		str := strings.Split(ins, "")
		for _, s := range str {
			i, _ := strconv.Atoi(s)
			tmp = append(tmp, Point{i, false})
		}
		grid = append(grid, tmp)
	}

	// drawGrid(grid)

	// Check grid.
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			// fmt.Println("checking:", grid[y][x].height)
			above, below, left, right := 9, 9, 9, 9

			// check x-1 / left
			if x != 0 {
				if grid[y][x].height > grid[y][x-1].height {
					continue
				}
				left = grid[y][x-1].height
			}
			// check x+1 / right
			if x != len(grid[y])-1 {
				if grid[y][x].height > grid[y][x+1].height {
					continue
				}
				right = grid[y][x+1].height
			}
			// check y-1 / above
			if y != 0 {
				if grid[y][x].height > grid[y-1][x].height {
					continue
				}
				above = grid[y-1][x].height
			}
			// check y+1 / below
			if y != len(grid)-1 {
				if grid[y][x].height > grid[y+1][x].height {
					continue
				}
				below = grid[y+1][x].height
			}

			if grid[y][x].height < above && grid[y][x].height < below && grid[y][x].height < left && grid[y][x].height < right {
				grid[y][x].lowest = true
				output += grid[y][x].height + 1
			}
		}
	}

	// drawGrid(grid)

	return output
}

func Part2(ctx context.Context, instructions []string) int {
	var (
		output int
	)

	return output
}

func drawGrid(grid [][]Point) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].lowest {
				fmt.Printf("%s%d%s", "\u001b[32m", grid[y][x].height, "\u001b[0m")
			} else {
				fmt.Print(grid[y][x].height)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
