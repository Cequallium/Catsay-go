package main

import (
	"fmt"
	c "github.com/fatih/color"
)

func main() {
	fmt.Printf("Usage: catsay %vhelo <Color (1 -> 8) >", c.New(c.FgBlue))
}
