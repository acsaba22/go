package main

import (
	"fmt"

	"github.com/acsaba22/go/conv"
)

func main() {
	h := conv.Meter(1.92) // Zurich - Stuttgart
	fmt.Printf("type: %T %v\n", h, h)
	hf := conv.MeterToFoot(h)
	fmt.Printf("type: %T %v\n", hf, hf)
}
