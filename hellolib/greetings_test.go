package hellolib_test

import (
	"fmt"

	"github.com/acsaba22/go/hellolib"
)

func ExampleGreeting() {
	fmt.Println(hellolib.Greeting())
	// Outputs:
	// Great, you used a lib. Hello btw.
}
