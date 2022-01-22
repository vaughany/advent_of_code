package aoc2021day3

import (
	"context"
	"testing"
)

func TestPart1(t *testing.T) {
	ctx := context.Background()
	instructions := []string{
		"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010",
	}

	result := Part1(ctx, instructions)
	expected := 198

	if result != expected {
		t.Errorf("Part1 was incorrect: got %d, want %d.", result, expected)
	}

}

func TestPart2(t *testing.T) {
	ctx := context.Background()
	instructions := []string{
		"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010",
	}

	result := Part2(ctx, instructions)
	expected := 230

	if result != expected {
		t.Errorf("Part2 was incorrect: got %d, want %d.", result, expected)
	}

}
