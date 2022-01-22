package aoc2021day16

import (
	"context"
	"fmt"
	"strconv"
)

type Packet struct {
	versionID    int64
	typeID       int64
	lengthTypeID string
	literal      string
}

// Part One:
func Part1(ctx context.Context, instructions string) int {
	var (
		output int64
	)

	instruction, err := strconv.ParseUint(instructions, 16, 32)
	if err != nil {
		fmt.Printf("%s", err)
	}
	// Convert int to binary representation
	// %024b indicates base 2, padding with 0, with 24 characters.
	bin := fmt.Sprintf("%024b", instruction)
	fmt.Println(bin)

	var packet Packet
	packet.versionID, _ = strconv.ParseInt(bin[0:3], 2, 64)
	packet.typeID, _ = strconv.ParseInt(bin[3:6], 2, 64)
	start := 6
	switch packet.typeID {
	case 4:
		for {
			fiveBits := bin[start : start+5]
			packet.literal += fiveBits[1:]
			fmt.Println(packet.literal)

			if fiveBits[0:1] == "0" {
				break
			}

			start += 5
		}
	default:
		packet.lengthTypeID = bin[6:7]
		switch packet.lengthTypeID {
		case "0":
			read, _ := strconv.ParseInt(bin[start:start+15], 2, 64)
			fmt.Println(read)
		case "1":
			read, _ := strconv.ParseInt(bin[start:start+11], 2, 64)
			fmt.Println(read)
		}

	}

	output, _ = strconv.ParseInt(packet.literal, 2, 64)
	fmt.Println(packet)

	return int(output)
}

// Part Two:
func Part2(ctx context.Context, instructions string) int {
	var (
		output int
	)

	return output
}
