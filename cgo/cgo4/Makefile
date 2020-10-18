all: main_bin

*.o: *.cpp
	g++ -Wall -c $^

libhello.a: hello.o
	ar rcs $@ $^

# -a is needed because it won't regenerate if .a changed.
main_bin: main.go libhello.a
	go build -a --ldflags "-linkmode external -extldflags -static" -o $@ $<

clean:
	$(RM) *.o *.a main_bin

.PHONY: all clean
