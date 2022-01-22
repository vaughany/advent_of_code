package aoc2021day17

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
)

type TargetArea struct {
	x1, x2, y1, y2 int
}

// Part One:
func Part1(ctx context.Context, instructions string) int {
	var (
		output int
	)

	targetArea := makeTargetArea(instructions)
	grid := makeGrid(targetArea)

	fmt.Printf("%v\n", targetArea)
	drawGrid(grid)

	return output
}

func makeTargetArea(ins string) TargetArea {
	re := regexp.MustCompile(`-?[0-9]{1,3}`)
	tmp := re.FindAllString(ins, 4)
	x1, _ := strconv.Atoi(tmp[0])
	x2, _ := strconv.Atoi(tmp[1])
	y1, _ := strconv.Atoi(tmp[2])
	y2, _ := strconv.Atoi(tmp[3])

	return TargetArea{x1, x2, y1, y2}
}

func makeGrid(targetArea TargetArea) [][]string {
	var (
		xx1, xx2, yy1, yy2 int
	)

	if targetArea.x1 > 0 {
		xx1 = 0
		xx2 = targetArea.x1
		if targetArea.x2 > xx2 {
			xx2 = targetArea.x2
		}
	}
	if targetArea.x1 < 0 {
		xx2 = 0
		xx1 = targetArea.x1
		if targetArea.x1 < xx1 {
			xx1 = targetArea.x1
		}
	}
	if targetArea.y1 > 0 {
		yy1 = 0
		yy2 = targetArea.y1
		if targetArea.y2 > yy2 {
			yy2 = targetArea.y2
		}
	}
	if targetArea.y1 < 0 {
		yy2 = 0
		yy1 = targetArea.y1
		if targetArea.y1 < yy1 {
			yy1 = targetArea.y1
		}
	}

	fmt.Println("XX and YY:", xx1, xx2, yy1, yy2)

	grid := make([][]string, 20)
	for i := range grid {
		grid[i] = make([]string, 40)
	}

	for j := 0; j < 20; j++ {
		for i := 0; i < 40; i++ {
			grid[j][i] = "."
		}
	}

	// Create target area.
	for j := 5; j <= 10; j++ {
		for i := 5; i <= 10; i++ {
			// fmt.Println("j", j, "i", i)
			grid[j][i] = "T"
		}
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
func Part2(ctx context.Context, instructions string) int {
	var (
		output int
	)

	return output
}
