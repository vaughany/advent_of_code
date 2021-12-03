package output

import (
	"fmt"
	"time"
)

var (
	title  = "\u001b[32m"
	info   = "\u001b[33m"
	timing = "\u001b[36m"
)

type TimeInfoType int

const (
	InfoTypeSetup TimeInfoType = iota
	InfoTypeOne
	InfoTypeTwo
	InfoTypeBoth
	InfoTypeEverything
)

func Title(year, day int) {
	sendToLog(title, fmt.Sprintf("Advent of Code %d, Day %d.", year, day))
}

func Info(text string) {
	sendToLog(info, text)
}

func TimeInfo(timeType TimeInfoType, timeDuration time.Duration) {
	var (
		output string
	)

	switch timeType {
	case InfoTypeSetup:
		output = "Setup took "
	case InfoTypeOne:
		output = "Part One took "
	case InfoTypeTwo:
		output = "Part Two took "
	case InfoTypeBoth:
		output = "Both Parts took "
	case InfoTypeEverything:
		output = "Everything took "
	}

	sendToLog(timing, fmt.Sprintf("%s%s", output, timeDuration))
}

func sendToLog(colour, contents string) {
	// log.Println(string(colour) + contents + string("\u001b[0m"))
	fmt.Println(string(colour) + contents + string("\u001b[0m"))
}

func Answer(part int, output interface{}) {
	partText := "One"

	if part == 2 {
		partText = "Two"
	}

	fmt.Printf("Part %s: %d\n", partText, output)
}
