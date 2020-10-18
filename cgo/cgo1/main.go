package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint() {
 printf("Hello world\n");
}
*/
import "C"

func main() {
	C.myprint()
}
