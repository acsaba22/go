1) Create a new test, and make it pass (rename package,...)

$ cp ../iowordfreq/iowordfreq_test.go ./fastWF_test.go


2) Make a new test which reads ../testdata/romeo_juliet.txt and checks that Romeo
   count is > 100.


3) Turn the previous into a banchmark test.

$ go test -run=NONE -bench=.


4) Start cpu profiling. Try at least these:

$ go test -run=NONE -bench=. -cpuprofile=cpu.out

$ go tool pprof text -nodecount=10 prof.test cpu.out
$ go tool pprof web prof.test cpu.out
$ go tool pprof prof.test cpu.out
(pprof) top10
(pprof) top20 -cum
(pprof) web
(pprof) list AddWords


5) Check mem profile.

$ go test -run=NONE -bench=. -benchmem
$ go test -run=NONE -bench=. -memprofile=mem.out
$ go tool pprof ...


6) Write a fastWF.go. Start with copying over the source code and try to
   optimeize it. Keep the original banchmark and create a new one so you can
   compare it.
