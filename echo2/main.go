package main

import (
	"flag"
	"fmt"
)

var s = flag.String("s", " ", "Separator between tokens")
var newLine = flag.Bool("nonewline", false, "Add new line at the end.")

func main() {
	flag.Parse()
	for i, a := range flag.Args() {
		if i != 0 {
			fmt.Printf("%s", *s)
		}
		fmt.Printf("%s", a)
	}
	if !*newLine {
		fmt.Println()
	}
}
