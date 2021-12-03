package loaders

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
)

// GetFilename creates the path and filename based on day and real / sample data.
func GetFilename(ctx context.Context, day int) string {
	var (
		suffix string
	)

	if ctx.Value("sample") == true {
		suffix = "-sample"
	}

	return fmt.Sprintf("inputs/%02d%s.txt", day, suffix)
}

// Read a file with many lines and return an array (of strings).
func LoadFile(filename string) []string {
	var (
		lines []string
	)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) < 1 {
		fmt.Println("Input file had no lines.")
		os.Exit(1)
	}

	// if debug {
	// 	info(fmt.Sprintf("Input file has %d lines / instructions.", len(lines)))
	// }

	return lines
}

func GetInputAsInts(filename string) []int {
	var (
		input  = LoadFile(filename)
		output []int
	)

	for _, line := range input {
		int, _ := strconv.Atoi(line)
		output = append(output, int)
	}

	return output
}

func GetInputAsStrings(filename string) []string {
	return LoadFile(filename)
}
