package main

import (
	"fmt"
	c "github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

func main() {
	// check strings for err
	args := os.Args[1:]
	if len(args) == 0 || strings.Contains(strings.ToLower(args[0]), "help") {
		helpMessage(0)
	}
	if len(args) != 2 {
		helpMessage(1)
	}

	phrase := args[0]
	color, err := strconv.Atoi(args[1])
	if err != nil || color > 8 {
		helpMessage(2)
	}
	if phrase == "" {
		helpMessage(3)
	}

	_, _ = Colorize(color).Printf("\n  /| ､\n (°､ ｡ 7  %v \n  |､  ~ヽ\n  じしf_,)〳\n", phrase)
}

func helpMessage(errType int) {
	defer os.Exit(0)
	var phrase = "Phrase"
	var color = "Color (1 to 8)"
	switch errType {
	case 1:
		phrase = c.RedString(phrase)
		color = c.RedString(color)
		fmt.Println(" Missing arguments")
	case 2:
		phrase = c.GreenString(phrase)
		color = c.RedString(color)
	default:

	}
	fmt.Printf(" Usage: catsay <%v> <%v>\n", phrase, color)
	fmt.Print(" Colors: ")
	_, _ = Colorize(9).Printf(" %d ", 1)
	for i := 10; i <= 16; i++ {
		_, _ = Colorize(i).Add(c.FgBlack).Printf(" %d ", i-8)
	}
	fmt.Println("")
}

func Colorize(color int) *c.Color {
	o := c.New()
	switch color {
	case 1:
		o = o.Add(c.FgBlack)
	case 2:
		o = o.Add(c.FgRed)
	case 3:
		o = o.Add(c.FgGreen)
	case 4:
		o = o.Add(c.FgYellow)
	case 5:
		o = o.Add(c.FgBlue)
	case 6:
		o = o.Add(c.FgMagenta)
	case 7:
		o = o.Add(c.FgCyan)
	case 8:
		o = o.Add(c.FgWhite)
	case 9:
		o = o.Add(c.BgBlack)
	case 10:
		o = o.Add(c.BgRed)
	case 11:
		o = o.Add(c.BgGreen)
	case 12:
		o = o.Add(c.BgYellow)
	case 13:
		o = o.Add(c.BgBlue)
	case 14:
		o = o.Add(c.BgMagenta)
	case 15:
		o = o.Add(c.BgCyan)
	case 16:
		o = o.Add(c.BgWhite)
	default:
		helpMessage(2)
	}
	return o
}
