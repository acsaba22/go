package hellolib_test

import (
	"fmt"

	"lib/hellolib"
)

func ExampleGreeting() {
	fmt.Println(hellolib.Greeting())
	// Outputs:
	// Great, you used a lib. Hello btw.
}
