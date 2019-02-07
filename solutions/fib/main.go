package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	fmt.Println("1:", fib(1))
	fmt.Println("5:", fib(5))
	fmt.Println("10:", fib(10))
}
