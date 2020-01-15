$ mkdir bin
$ cd bin
$ go tool compile -help
  -L    show full file names in error messages
  -N    disable optimizations
  -l    disable inlining
  -m    print optimization decisions
$ go build -gcflags '-N -L -l -m' -o m.exe mycompany.com/firstgo/typedebugmain
$ ls
m.exe
$ go tool compile -N -L -l -m ../typedebug/typedebug.go
$ ls
m.exe  typedebug3.o
$ go tool objdump typedebug3.o 
TEXT %22%22.SliceIsItNil(SB) gofile../<...>/typedebug.go
  typedebug3.go:8       0x7fb                   c644242000              MOVB $0x0, 0x20(SP)
  typedebug3.go:9       0x800                   48837c240800            CMPQ $0x0, 0x8(SP)
  typedebug3.go:9       0x806                   0f94442420              SETE 0x20(SP)
  typedebug3.go:9       0x80b                   c3                      RET

...
$ go tool objdump m.exe | less

Let's look at the results together