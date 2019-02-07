// Package hellolib is the first example package.
// Usage:
//
// import (
//   "fmt"
//
//   "github.com/acsaba22/go/hellolib"
// )
//
// func main() {
// 	fmt.Println(hellolib.Greeting())
// }
package hellolib

func actualGreeting() string {
	return "Hello btw."
}

// Greeting is an exported function, feel free to use it.
// For usage example see the package level documentation.
func Greeting() string {
	return "Great, you used a lib. " + actualGreeting()
}
