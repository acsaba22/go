package main

// go build main.go
// ldd ./main

//#include "hello.c"
import "C"

func main() {
	C.myprint()
}
