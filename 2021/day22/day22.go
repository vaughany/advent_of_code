package aoc2021day22

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CoordIns struct {
	onOrOff                bool
	x1, x2, y1, y2, z1, z2 int
}

// Part One:
func Part1(ctx context.Context, instructions []string) int {
	var (
		coordsIns []CoordIns
		coordsOn  int
	)

	grid := makeGrid()

	// fmt.Println("0:", grid[0][0][0])
	// fmt.Println("1:", grid[1][1][1])

	re1 := regexp.MustCompile(`(-?[0-9]{1,6})`)
	for _, instruction := range instructions {
		coords := CoordIns{}
		if strings.Contains(instruction, "on") {
			coords.onOrOff = true
		}

		tmp := re1.FindAllString(instruction, 6)
		fmt.Println(tmp)
		coords.x1, _ = strconv.Atoi(tmp[0])
		coords.x2, _ = strconv.Atoi(tmp[1])
		coords.y1, _ = strconv.Atoi(tmp[2])
		coords.y2, _ = strconv.Atoi(tmp[3])
		coords.z1, _ = strconv.Atoi(tmp[4])
		coords.z2, _ = strconv.Atoi(tmp[5])
		fmt.Println("coords:", coords)

		coordsIns = append(coordsIns, coords)
	}

	for _, coord := range coordsIns {
		grid, coordsOn = flipGrid(grid, coordsOn, coord)
		fmt.Println("coordsOn:", countGrid(grid))
	}

	return coordsOn
}

func flipGrid(grid [][][]bool, coordsOn int, coord CoordIns) ([][][]bool, int) {

	for zz := coord.z1; zz <= coord.z2; zz++ {
		for yy := coord.y1; yy <= coord.y2; yy++ {
			for xx := coord.x1; xx <= coord.x2; xx++ {
				grid[zz][yy][xx] = coord.onOrOff
			}
		}
	}

	return grid, coordsOn
}

func countGrid(grid [][][]bool) int {
	var count int

	for z := 0; z < len(grid); z++ {
		for y := 0; y < len(grid[z]); y++ {
			for x := 0; x < len(grid[z][y]); x++ {
				if grid[z][y][x] {
					count++
				}
			}
		}
	}

	return count
}

func makeGrid() [][][]bool {
	grid := make([][][]bool, 50)

	for i := range grid {
		grid[i] = make([][]bool, 50)
		for j := range grid[i] {
			grid[i][j] = make([]bool, 50)
		}
	}

	return grid
}

// Part Two:
func Part2(ctx context.Context, instructions []string) int {
	var (
		output int
	)

	return output
}
