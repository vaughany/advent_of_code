package main

import (
	"context"
	"flag"

	aoc2021day1 "github.com/vaughany/advent_of_code/day01"
	aoc2021day2 "github.com/vaughany/advent_of_code/day02"
	aoc2021day3 "github.com/vaughany/advent_of_code/day03"
	"github.com/vaughany/advent_of_code/loaders"
	"github.com/vaughany/advent_of_code/output"
)

func main() {

	var (
		debug bool
		// timing bool
		day int
		// filename string
	)

	// flag.StringVar(&filename, "f", filename, "Specify a file to read input from")
	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	// flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.IntVar(&day, "day", day, "Run just this day")
	flag.Parse()

	ctx := context.WithValue(context.Background(), "debug", debug)
	// ctx = context.WithValue(ctx, "timing", timing)

	// 2021, day 1.
	if day == 0 || day == 1 {
		output.Title(2021, 1)
		d1ins := loaders.GetInputAsInts(loaders.GetFilename(1))
		output.Answer(1, aoc2021day1.Part1(ctx, d1ins))
		output.Answer(2, aoc2021day1.Part2(ctx, d1ins))
	}

	// 2021, day 2.
	if day == 0 || day == 2 {
		output.Title(2021, 2)
		d2ins := loaders.GetInputAsStrings(loaders.GetFilename(2))
		output.Answer(1, aoc2021day2.Part1(ctx, d2ins))
		output.Answer(2, aoc2021day2.Part2(ctx, d2ins))
	}

	// 2021, day 3.
	if day == 0 || day == 3 {
		output.Title(2021, 3)
		d3ins := loaders.GetInputAsStrings(loaders.GetFilename(3))
		output.Answer(1, aoc2021day3.Part1(ctx, d3ins))
		output.Answer(2, aoc2021day3.Part2(ctx, d3ins))
	}
}
