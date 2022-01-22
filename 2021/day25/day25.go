package aoc2021day25

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	empty = "."
	down  = "v"
	right = ">"
)

// Part One:
func Part1(ctx context.Context, instructions []string) int {
	var (
		debug = ctx.Value("debug").(bool)
		maxY  = len(instructions)
		maxX  = len(instructions[0])
	)

	fmt.Println("maxX:", maxX, "maxY:", maxY)

	// Make the grid
	grid := populateGrid(makeGrid(maxX, maxY), instructions)
	if debug {
		drawGrid(grid)
		fmt.Printf("%#v\n", grid)
	}

	// Process right.
	loops := 1
	for {
		newGrid := makeGrid(maxX, maxY)

		// Right
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == right {
					if canMoveRight(grid, x, y, maxX) {
						newGrid = moveRight(newGrid, x, y, maxX)
						x++
					} else {
						newGrid[y][x] = right
					}
				} else if grid[y][x] == down {
					newGrid[y][x] = down
				} else if grid[y][x] == empty {
					newGrid[y][x] = empty
				}
			}
		}

		// Down.
		for x := 0; x < len(grid[0]); x++ {
			for y := 0; y < len(grid); y++ {
				if grid[y][x] == down {
					if canMoveDown(grid, x, y, maxY) {
						newGrid = moveDown(newGrid, x, y, maxY)
						y++
					} else {
						newGrid[y][x] = down
					}
				} else if grid[y][x] == right {
					newGrid[y][x] = right
				} else if grid[y][x] == empty {
					newGrid[y][x] = empty
				}
			}
		}

		grid = newGrid
		fmt.Println("loop:", loops)
		drawGrid(grid)
		loops++
		if loops > 1 {
			break
		}

		time.Sleep(time.Second)
	}

	return 0
}

func canMoveRight(grid [][]string, xx, yy, maxX int) bool {
	xx++
	if xx == maxX {
		xx = 0
	}

	return grid[yy][xx] == empty
}

func moveRight(grid [][]string, xx, yy, maxX int) [][]string {
	xx2 := xx + 1
	if xx2 == maxX {
		xx2 = 0
	}

	grid[yy][xx], grid[yy][xx2] = empty, right

	return grid
}

func canMoveDown(grid [][]string, xx, yy, maxY int) bool {
	yy++
	if yy == maxY {
		yy = 0
	}

	return grid[yy][xx] == empty
}

func moveDown(grid [][]string, xx, yy, maxY int) [][]string {
	yy2 := yy + 1
	if yy2 == maxY {
		yy2 = 0
	}

	grid[yy][xx], grid[yy2][xx] = empty, down

	return grid
}

func makeGrid(xx, yy int) [][]string {
	grid := make([][]string, yy)
	for i := range grid {
		grid[i] = make([]string, xx)
	}

	return grid
}

func populateGrid(grid [][]string, instructions []string) [][]string {
	for y := 0; y < len(instructions); y++ {
		grid[y] = strings.Split(instructions[y], "")
	}

	return grid
}

func drawGrid(grid [][]string) {
	var out string

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			out += grid[y][x]
		}
		out += "\n"
	}
	out += "\n"

	fmt.Print(out)
}

// Part Two:
func Part2(ctx context.Context, instructions []string) int {
	return 0
}
