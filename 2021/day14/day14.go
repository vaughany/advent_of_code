package aoc2021day14

import (
	"context"
	"fmt"
	"math"
	"strings"
)

// Part One:
func Part1(ctx context.Context, instructions []string) int {
	var (
		debug = ctx.Value("debug").(bool)
		rules = make(map[string]string)
	)

	template := instructions[0]

	for i := 2; i <= len(instructions)-1; i++ {
		tmp := strings.Split(instructions[i], " -> ")
		rules[tmp[0]] = tmp[1]
	}

	if debug {
		fmt.Println("rules:", rules)
	}

	for j := 1; j <= 10; j++ {
		if debug {
			fmt.Println("loop:", j)
			fmt.Println("template:", template)

		}

		var newTemplate string
		for i := 0; i <= len(template)-2; i++ {
			if rule, ok := rules[template[i:i+2]]; ok {
				newTemplate += string(template[i]) + rule
			} else {
				newTemplate += string(template[i])
			}
		}
		template = newTemplate + template[len(template)-1:]
	}

	// Sorting
	polymerCount := make(map[string]int)
	for _, polymer := range strings.Split(template, "") {
		p := string(polymer)
		polymerCount[p]++
	}

	if debug {
		fmt.Println(polymerCount)
	}

	min, max := math.MaxInt, 0
	for _, i := range polymerCount {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	return max - min
}

// Part Two:
func Part2(ctx context.Context, instructions []string) int {
	var (
		debug = ctx.Value("debug").(bool)
		rules = make(map[string]string)
	)

	template := instructions[0]

	for i := 2; i <= len(instructions)-1; i++ {
		tmp := strings.Split(instructions[i], " -> ")
		rules[tmp[0]] = tmp[1]
	}

	if debug {
		fmt.Println("rules:", rules)
	}

	for j := 1; j <= 10; j++ {
		if debug {
			fmt.Println("loop:", j)
			fmt.Println("template:", template)

		}

		var newTemplate string
		for i := 0; i <= len(template)-2; i++ {
			if rule, ok := rules[template[i:i+2]]; ok {
				newTemplate += string(template[i]) + rule
			} else {
				newTemplate += string(template[i])
			}
		}
		template = newTemplate + template[len(template)-1:]
	}

	// Sorting
	polymerCount := make(map[string]int)
	for _, polymer := range strings.Split(template, "") {
		p := string(polymer)
		polymerCount[p]++
	}

	if debug {
		fmt.Println(polymerCount)
	}

	min, max := math.MaxInt, 0
	for _, i := range polymerCount {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	return max - min
}
