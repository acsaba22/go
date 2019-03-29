1) edit server.go, add profiling to the import:

import _ "net/http/pprof"

2) run the server

3) check out the page

http://localhost:8000/count?n=10000000000

4) check out profiling page:
http://localhost:8000/debug/pprof/

5) Find the bottleneck with the go tool:

$ go tool pprof http://localhost:8000/debug/pprof/profile?seconds=30s
