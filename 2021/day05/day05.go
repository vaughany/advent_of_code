package aoc2021day5

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
)

type Line struct {
	x1, y1, x2, y2 int
}

// Part One: 5294
// Part Two: 21698
// Upping the Ante: 1967556606 https://www.reddit.com/r/adventofcode/comments/r9hpfs/2021_day_5_bigger_vents/
func Part1And2(ctx context.Context, instructions []string) (int, int) {
	var (
		lines []Line
		debug = ctx.Value("debug").(bool)
		size  int
	)

	re := regexp.MustCompile(`[0-9]{1,9}`)
	for _, instruction := range instructions {
		coords := re.FindAll([]byte(instruction), 4)
		if len(coords) != 4 {
			panic("Error: four items not found in a slice.")
		}
		x1, _ := strconv.Atoi(string(coords[0]))
		y1, _ := strconv.Atoi(string(coords[1]))
		x2, _ := strconv.Atoi(string(coords[2]))
		y2, _ := strconv.Atoi(string(coords[3]))

		lines = append(lines, Line{x1, y1, x2, y2})

		if x2 > size {
			size = x2
		}
		if y2 > size {
			size = y2
		}
	}
	// Add one more on.
	size++
	if debug {
		fmt.Println("Grid height and width:", size)
	}

	if debug {
		fmt.Println(lines)
	}

	// Create and initialise grid.
	grid1 := make([][]int, size)
	grid2 := make([][]int, size)
	for i := range grid1 {
		grid1[i] = make([]int, size)
		grid2[i] = make([]int, size)
	}

	if debug {
		drawGrid(size, grid1)
	}

	for _, line := range lines {
		// Part one only uses horizontal or vertical lines.
		if line.x1 == line.x2 || line.y1 == line.y2 {
			grid1 = drawLine(grid1, line.x1, line.y1, line.x2, line.y2)
		}
		// Part two: go nuts!
		grid2 = drawLine(grid2, line.x1, line.y1, line.x2, line.y2)
	}

	if debug {
		drawGrid(size, grid1)
	}

	return countGrid(size, grid1), countGrid(size, grid2)
}

func countGrid(size int, grid [][]int) int {
	var twoOrMore int

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] > 1 {
				twoOrMore++
			}
		}
	}

	return twoOrMore
}

func drawLine(grid [][]int, x1, y1, x2, y2 int) [][]int {
	// fmt.Println(x1, y1, x2, y2)

	switch {
	case x1 == x2:
		// fmt.Println("x1 == x2")
		if y1 < y2 {
			for y := y1; y <= y2; y++ {
				// fmt.Printf("%d,%d;", x1, y)
				grid[y][x1]++
			}
		} else {
			for y := y2; y <= y1; y++ {
				// fmt.Printf("%d,%d;", x1, y)
				grid[y][x1]++
			}
		}
	case y1 == y2:
		// fmt.Println("y1 == y2")
		if x1 < x2 {
			for x := x1; x <= x2; x++ {
				// fmt.Printf("%d,%d;", x, y1)
				grid[y1][x]++
			}
		} else {
			for x := x2; x <= x1; x++ {
				// fmt.Printf("%d,%d;", x, y1)
				grid[y1][x]++
			}
		}
	default:
		// fmt.Println("Diagonal")
		if x1 < x2 && y1 < y2 {
			for {
				// fmt.Printf("%d,%d;", x1, y1)
				grid[y1][x1]++
				if x1 == x2 {
					break
				}
				x1++
				y1++
			}
		}
		if x1 > x2 && y1 > y2 {
			for {
				// fmt.Printf("%d,%d;", x1, y1)
				grid[y1][x1]++
				if x1 == x2 {
					break
				}
				x1--
				y1--
			}
		}
		if x1 < x2 && y1 > y2 {
			for {
				// fmt.Printf("%d,%d;", x1, y1)
				grid[y1][x1]++
				if x1 == x2 {
					break
				}
				x1++
				y1--
			}
		}
		if x1 > x2 && y1 < y2 {
			for {
				// fmt.Printf("%d,%d;", x1, y1)
				grid[y1][x1]++
				if x1 == x2 {
					break
				}
				x1--
				y1++
			}
		}
	}

	// fmt.Println()
	return grid
}

// Looks fine for the sample instructions, which are 10x10 maximum.
// TODO: consider outputting to a png, or animated gif or something.
func drawGrid(size int, grid [][]int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(grid[y][x])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
