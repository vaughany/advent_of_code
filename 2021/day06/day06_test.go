package aoc2021day6

import "testing"

// https://blog.alexellis.io/golang-writing-unit-tests/
// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

var (
	input           = []int{4, 1, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 3, 4, 1, 1, 1, 3, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 3, 1, 1, 1, 5, 1, 2, 1, 1, 5, 3, 4, 2, 1, 1, 4, 1, 1, 5, 1, 1, 5, 5, 1, 1, 5, 2, 1, 4, 1, 2, 1, 4, 5, 4, 1, 1, 1, 1, 3, 1, 1, 1, 4, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 4, 4, 1, 1, 3, 1, 3, 2, 4, 3, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 2, 5, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 4, 1, 5, 1, 3, 1, 1, 1, 1, 1, 5, 1, 1, 1, 3, 1, 2, 1, 2, 1, 3, 4, 5, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 3, 1, 1, 4, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 3, 2, 1, 1, 1, 4, 2, 1, 1, 1, 4, 1, 1, 2, 3, 1, 4, 1, 5, 1, 1, 1, 2, 1, 5, 3, 3, 3, 1, 5, 3, 1, 1, 1, 1, 1, 1, 1, 1, 4, 5, 3, 1, 1, 5, 1, 1, 1, 4, 1, 1, 5, 1, 2, 3, 4, 2, 1, 5, 2, 1, 2, 5, 1, 1, 1, 1, 4, 1, 2, 1, 1, 1, 2, 5, 1, 1, 5, 1, 1, 1, 3, 2, 4, 1, 3, 1, 1, 2, 1, 5, 1, 3, 4, 4, 2, 2, 1, 1, 1, 1, 5, 1, 5, 2}
	days1           = 80
	days2           = 256
	expectedResult1 = 386640
	expectedResult2 = 1733403626279
)

func TestFigureOutFishNumbers(t *testing.T) {
	result := figureOutFishNumbers(input, days1)
	if result != expectedResult1 {
		t.Errorf("figureOutFishNumbers was incorrect, got: %d, want: %d.", result, expectedResult1)
	}

	result = figureOutFishNumbers(input, days2)
	if result != expectedResult2 {
		t.Errorf("figureOutFishNumbers was incorrect, got: %d, want: %d.", result, expectedResult2)
	}
}

func BenchmarkFigureOutFishNumbers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		figureOutFishNumbers(input, days2)
	}
}
