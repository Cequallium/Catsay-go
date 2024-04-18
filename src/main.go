package main

import (
	"fmt"
	c "github.com/fatih/color"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		helpMessage(1)
	}
	phrase := args[0]
	color, err := strconv.Atoi(args[1])
	if err != nil {
		panic("INVALID COLOR")
	}
	if phrase == "" {
		panic("MISSING PHRASE")
	}

	_, _ = Colorize(color).Printf("  /| ､\n (°､ ｡ 7  %v \n  |､  ~ヽ\n  じしf_,)〳\n", phrase)
}

func helpMessage(errType int) {
	var phrase = "<Phrase>"
	var color = "<Color (1 -> 8)>"
	switch errType {
	case 1:
		phrase = c.RedString(phrase)
		color = c.RedString(color)
	case 2:
		color = c.RedString(color)
	case 3:
		phrase = c.RedString(phrase)
	}
	fmt.Printf("Usage: catsay %v %v\n")
	os.Exit(0)
}

func Colorize(color int) *c.Color {
	outcolor := c.New()
	switch color {
	case 1:
		outcolor = outcolor.Add(c.FgBlack)
	case 2:
		outcolor = outcolor.Add(c.FgRed)
	case 3:
		outcolor = outcolor.Add(c.FgGreen)
	case 4:
		outcolor = outcolor.Add(c.FgYellow)
	case 5:
		outcolor = outcolor.Add(c.FgBlue)
	case 6:
		outcolor = outcolor.Add(c.FgMagenta)
	case 7:
		outcolor = outcolor.Add(c.FgCyan)
	case 8:
		outcolor = outcolor.Add(c.FgWhite)
	default:
		panic("COLOR MUST BE <=8")
	}
	return outcolor
}
