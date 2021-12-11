package aoc2021day8

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Part One: 532
func Part1(ctx context.Context, instructions []string) int {
	var (
		output int
	)

	for _, ins := range instructions {
		someStrings := strings.Split(ins, " | ")
		moreStrings := strings.Split(someStrings[1], " ")

		var one, four, seven, eight int
		for _, str := range moreStrings {
			switch len(str) {
			case 2:
				one++
			case 4:
				four++
			case 3:
				seven++
			case 7:
				eight++
			}
		}
		output += one + four + seven + eight
	}

	return output
}

// Part Two: 1011284
func Part2(ctx context.Context, instructions []string) int {
	var (
		// debug  = ctx.Value("debug").(bool)
		part2  int
		digits = make(map[string]string)
	)

	// Two steps:
	//   1. figure out the frequencies of each character over all input values (left of '|')
	//   2. replace each character with its frequency and sort the result (e.g. acedgfb becomes 4677889)
	// Each input/output value can then be decoded using a pre-computed hash map.

	// Frequency of each segment:
	// a: 8    aaaa     8888
	// b: 6   b    c   6    8
	// c: 8   b    c   6    8
	// d: 7    dddd     7777
	// e: 4   e    f   4    9
	// f: 9   e    f   4    9
	// g: 7    gggg     7777

	// Frequency of the segment, sorted.
	digits["467889"] = "0"  // abcefg
	digits["89"] = "1"      // cf
	digits["47788"] = "2"   // acdeg
	digits["77889"] = "3"   // acdfg
	digits["6789"] = "4"    // bcdf
	digits["67789"] = "5"   // abdfg
	digits["467789"] = "6"  // abdefg
	digits["889"] = "7"     // ace
	digits["4677889"] = "8" // abcdefg
	digits["677889"] = "9"  // abcdfg

	for _, ins := range instructions {
		someStrings := strings.Split(ins, " | ")
		input := someStrings[0]
		output := someStrings[1]

		// Discern the frequency of the letters.
		var segment = make(map[string]int)
		for _, letter := range []string{"a", "b", "c", "d", "e", "f", "g"} {
			segment[letter] = strings.Count(input, letter)
		}

		var newout string
		for _, str := range strings.Split(output, " ") {
			var out string
			for _, s := range str {
				out += fmt.Sprintf("%d", segment[string(s)])
			}
			out = sortString(out)
			newout += digits[out]
		}

		tmp, _ := strconv.Atoi(newout)
		part2 += tmp
	}

	return part2
}

func sortString(in string) string {
	out := strings.Split(in, "")
	sort.Strings(out)
	return strings.Join(out, "")
}
