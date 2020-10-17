package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	for _, v := range os.Args[1:] {
		fmt.Println(v)
	}
}
