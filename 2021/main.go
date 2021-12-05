package main

import (
	"context"
	"flag"
	"time"

	aoc2021day1 "github.com/vaughany/advent_of_code/day01"
	aoc2021day2 "github.com/vaughany/advent_of_code/day02"
	aoc2021day3 "github.com/vaughany/advent_of_code/day03"
	aoc2021day4 "github.com/vaughany/advent_of_code/day04"
	aoc2021day5 "github.com/vaughany/advent_of_code/day05"
	"github.com/vaughany/advent_of_code/loaders"
	"github.com/vaughany/advent_of_code/output"
)

func main() {

	var (
		debug  bool
		timing bool
		sample bool
		day    int

		timeSetup time.Time
		timeOne   time.Time
		timeTwo   time.Time
	)

	flag.BoolVar(&debug, "d", debug, "Display debugging information")
	flag.BoolVar(&timing, "t", timing, "Display timing information")
	flag.IntVar(&day, "day", day, "Run just this day")
	flag.BoolVar(&sample, "s", sample, "Use the sample input data")
	flag.Parse()

	ctx := context.WithValue(context.Background(), "debug", debug)
	ctx = context.WithValue(ctx, "sample", sample)

	// 2021, day 1.
	if day == 0 || day == 1 {
		output.Title(2021, 1)
		ins := loaders.GetInputAsInts(loaders.GetFilename(ctx, 1))
		output.Answer(1, aoc2021day1.Part1(ctx, ins))
		output.Answer(2, aoc2021day1.Part2(ctx, ins))
	}

	// 2021, day 2.
	if day == 0 || day == 2 {
		output.Title(2021, 2)
		ins := loaders.GetInputAsStrings(loaders.GetFilename(ctx, 2))
		output.Answer(1, aoc2021day2.Part1(ctx, ins))
		output.Answer(2, aoc2021day2.Part2(ctx, ins))
	}

	// 2021, day 3.
	if day == 0 || day == 3 {
		timeSetup = time.Now()
		output.Title(2021, 3)
		ins := loaders.GetInputAsStrings(loaders.GetFilename(ctx, 3))
		if timing {
			output.TimeInfo(output.InfoTypeSetup, time.Since(timeSetup))
			timeOne = time.Now()
		}

		output.Answer(1, aoc2021day3.Part1(ctx, ins))
		if timing {
			output.TimeInfo(output.InfoTypeOne, time.Since(timeOne))
			timeTwo = time.Now()
		}

		output.Answer(2, aoc2021day3.Part2(ctx, ins))
		if timing {
			output.TimeInfo(output.InfoTypeTwo, time.Since(timeTwo))
			output.TimeInfo(output.InfoTypeBoth, time.Since(timeOne))
			output.TimeInfo(output.InfoTypeEverything, time.Since(timeSetup))
		}
	}

	// 2021, day 4.
	if day == 0 || day == 4 {
		timeSetup = time.Now()
		output.Title(2021, 4)
		ins := loaders.GetInputAsStrings(loaders.GetFilename(ctx, 4))
		if timing {
			output.TimeInfo(output.InfoTypeSetup, time.Since(timeSetup))
			timeOne = time.Now()
		}

		output.Answer(1, aoc2021day4.Part1(ctx, ins))
		if timing {
			output.TimeInfo(output.InfoTypeOne, time.Since(timeOne))
			timeTwo = time.Now()
		}

		output.Answer(2, aoc2021day4.Part2(ctx, ins))
		if timing {
			output.TimeInfo(output.InfoTypeTwo, time.Since(timeTwo))
			output.TimeInfo(output.InfoTypeBoth, time.Since(timeOne))
			output.TimeInfo(output.InfoTypeEverything, time.Since(timeSetup))
		}
	}

	// 2021, day 5.
	if day == 0 || day == 5 {
		timeSetup = time.Now()
		output.Title(2021, 5)
		ins := loaders.GetInputAsStrings(loaders.GetFilename(ctx, 5))
		if timing {
			output.TimeInfo(output.InfoTypeSetup, time.Since(timeSetup))
			timeOne = time.Now()
		}

		part1, part2 := aoc2021day5.Part1And2(ctx, ins)
		output.Answer(1, part1)
		output.Answer(2, part2)
		if timing {
			output.TimeInfo(output.InfoTypeBoth, time.Since(timeOne))
			output.TimeInfo(output.InfoTypeEverything, time.Since(timeSetup))
		}
	}
}
