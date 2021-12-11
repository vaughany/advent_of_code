package aoc2021day10

import (
	"context"
	"strings"
)

// Part One:
func Part1(ctx context.Context, instructions []string) int {
	var (
		illegals string
	)

	for _, ins := range instructions {
		var chunk string
		var done bool
		// fmt.Println("INS:", ins)
		for _, i := range strings.Split(ins, "") {
			if i == "(" || i == "[" || i == "{" || i == "<" {
				chunk += i
			} else {
				// Determine if the character is 'legit' or not.
				switch i {
				case ")":
					if getLastChar(chunk) == "(" {
						chunk = removeLastChar(chunk)
					} else {
						illegals += i
						// fmt.Println("ILLEGAL", i)
						done = true
					}
				case "]":
					if getLastChar(chunk) == "[" {
						chunk = removeLastChar(chunk)
					} else {
						illegals += i
						// fmt.Println("ILLEGAL", i)
						done = true
					}
				case "}":
					if getLastChar(chunk) == "{" {
						chunk = removeLastChar(chunk)
					} else {
						illegals += i
						// fmt.Println("ILLEGAL", i)
						done = true
					}
				case ">":
					if getLastChar(chunk) == "<" {
						chunk = removeLastChar(chunk)
					} else {
						illegals += i
						// fmt.Println("ILLEGAL", i)
						done = true
					}
				}
			}
			// fmt.Println(index, "-", chunk)
			if done {
				break
			}
		}
		if done {
			done = false
			continue
		}
	}

	// fmt.Println("illegals:", illegals)

	return getScore(illegals)
}

func getLastChar(in string) string {
	if len(in) == 0 {
		return ""
	}
	return in[len(in)-1:]
}

func removeLastChar(in string) string {
	// if len(in) == 0 {
	// return ""
	// }
	return in[:len(in)-1]
}

func getScore(in string) int {
	var out int

	for _, i := range in {
		switch string(i) {
		case ")":
			out += 3
		case "]":
			out += 57
		case "}":
			out += 1197
		case ">":
			out += 25137
		}
	}

	return out
}

func Part2(ctx context.Context, instructions []string) int {
	var (
		output int
	)

	return output
}
