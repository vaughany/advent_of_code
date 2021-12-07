package aoc2021day7

import (
	"context"
	"testing"
)

var (
	input           = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expectedResult1 = 37
	expectedResult2 = 168
)

func TestPart1And2(t *testing.T) {

	result1, result2 := Part1And2(context.Background(), input)

	if result1 != expectedResult1 {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", result1, expectedResult1)
	}
	if result2 != expectedResult2 {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", result2, expectedResult2)
	}
}

func BenchmarkPart1And2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1And2(context.Background(), input)
	}
}
