package main

// g++ -Wall -c hello.cpp
// ar rcs libhello.a hello.o
// go build -a --ldflags "-linkmode external -extldflags -static" -o main_bin3 main.go
// ldd main_bin

// objdump -x hello.o | less
// Contains myprint
// ar tOv libhello.a | less

/*
#cgo LDFLAGS: -L./ -lhello -lstdc++
#include "hello.h"
*/
import "C"

func main() {
	C.MyPrint()
	C.MyPrintCPP()
}
