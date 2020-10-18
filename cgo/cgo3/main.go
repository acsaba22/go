package main

// go build main.go
// ldd ./main

//#include "hello.cpp"
import "C"

func main() {
	C.myprint()
}
