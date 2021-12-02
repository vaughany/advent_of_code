package output

import (
	"fmt"
)

var (
	title  = "\u001b[32m"
	info   = "\u001b[33m"
	timing = "\u001b[36m"
)

func Title(year, day int) {
	sendToLog(title, fmt.Sprintf("Advent of Code %d, Day %d.", year, day))
}

func Info(text string) {
	sendToLog(info, text)
}

func TimeInfo(text string) {
	sendToLog(timing, text)
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
