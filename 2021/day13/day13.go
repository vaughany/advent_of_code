package aoc2021day13

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
)

type Dot struct {
	x, y int
}

type Fold struct {
	dir   string
	coord int
}

const (
	transparent = " "
	dotted      = "#"
	foldHoriz   = "="
	foldVert    = "|"
)

// Part One: 850
// Part Two: AHGCPGAU
func Part1And2(ctx context.Context, instructions []string) (int, string) {
	var (
		debug      = ctx.Value("debug").(bool)
		partOne    int
		partTwo    string
		coordsIns  []string
		foldsIns   []string
		coords     []Dot
		folds      []Fold
		maxX, maxY int
	)

	var fold bool
	for _, instruction := range instructions {
		if instruction == "" {
			fold = true
			continue
		}

		switch fold {
		case false:
			coordsIns = append(coordsIns, instruction)
		case true:
			foldsIns = append(foldsIns, instruction)
		}
	}

	re1 := regexp.MustCompile(`([0-9]{1,4})`)
	for _, coord := range coordsIns {
		tmp := re1.FindAllString(coord, 2)
		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])
		coords = append(coords, Dot{x, y})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	if debug {
		fmt.Println("COORDS:", coords)
		fmt.Println("MAX X:", maxX, "MAX Y:", maxY)
	}

	re2 := regexp.MustCompile(`([xy])=([0-9]{1,3})`)
	for _, fold := range foldsIns {
		tmp := re2.FindStringSubmatch(fold)
		dir := tmp[1]
		coord, _ := strconv.Atoi(string(tmp[2]))
		folds = append(folds, Fold{dir, coord})
	}
	if debug {
		fmt.Println("FOLDS:", folds)
	}

	// Make the grid
	grid := makeGrid(coords, maxX, maxY)
	if debug {
		fmt.Println(drawGrid(grid))
	}

	// Fold one and count, for part one.
	grid = foldGrid(grid, folds[0])
	if debug {
		fmt.Println(drawGrid(grid))
	}
	partOne = countGrid(grid)

	// Complete the rest of the folds for part two.
	for i := 1; i < len(folds); i++ {
		grid = foldGrid(grid, folds[i])
		if debug {
			fmt.Println(drawGrid(grid))
		}
	}

	if debug {
		partTwo = drawGrid(grid)
	} else {
		partTwo = drawFinalGrid(grid)
	}

	return partOne, partTwo
}

func foldGrid(grid [][]string, fold Fold) [][]string {
	if fold.dir == "y" {
		for i := 0; i <= fold.coord; i++ {
			for x := 0; x < len(grid[fold.coord]); x++ {
				if grid[fold.coord+i][x] == dotted {
					grid[fold.coord-i][x] = dotted
					// Remove the dot from the grid, instead of resizing the grid.
					grid[fold.coord+i][x] = transparent
				}
				grid[fold.coord][x] = foldHoriz
			}
		}
	} else {
		for i := 0; i <= fold.coord; i++ {
			for y := 0; y < len(grid); y++ {
				if grid[y][fold.coord+i] == dotted {
					grid[y][fold.coord-i] = dotted
					// Remove the dot from the grid, instead of resizing the grid.
					grid[y][fold.coord+i] = transparent
				}
				grid[y][fold.coord] = foldVert
			}
		}
	}

	return grid
}

func makeGrid(dots []Dot, xx, yy int) [][]string {
	// Make the grid.
	grid := make([][]string, yy+1)
	for i := range grid {
		grid[i] = make([]string, xx+1)
	}

	// Populate it with... something.
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x] = transparent
		}
	}

	// Add in the dots.
	for _, dot := range dots {
		grid[dot.y][dot.x] = dotted
	}

	return grid
}

func drawGrid(grid [][]string) string {
	var out string

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			out += grid[y][x]
		}
		out += "\n"
	}
	out += "\n"

	return out
}

func drawFinalGrid(grid [][]string) string {
	var out string

	for y := 0; y < 6; y++ {
		for x := 0; x < 39; x++ {
			out += grid[y][x]
		}
		out += "\n"
	}

	return out[:len(out)-2]
}

func countGrid(grid [][]string) int {
	var count int

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == dotted {
				count++
			}
		}
	}

	return count
}
