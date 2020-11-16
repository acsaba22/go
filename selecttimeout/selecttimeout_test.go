package main_test

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"testing"
	"time"
)

// This file goes through the explanation of why channels time out even when there is a value right away.
// Read through the comments of the functions to have an understanding of why I think my explanation is correct.
//
// TestOriginal           - the original situation we want to explain.
// TestOriginalHistogram  - has my explanation of what is actually happening.
// TestFixedHistogram     - has my proposed solution. Your solution is also fine IMHO.

const million = 1000 * 1000
const count = 10 * million // every test will do this many iterations.

// This is the original setup thanks for providing a concise test case.
//
// Even though we add a value to the channel, we don't get it back sometimes.
// How is this possible?
// Note that we gave a timeout of 100 Microsecond = 100,000 Nanosecond
//
// Here is a run:
// $ go test -p=1 -v -run=Original$ hello/selecttimeout_test.go
// === RUN   TestOriginal
// channel not read with timeout! 298227
// channel not read with timeout! 1031706
// channel not read with timeout! 1122213
// channel not read with timeout! 1660480
// channel not read with timeout! 2070965
// channel not read with timeout! 2444345
// channel not read with timeout! 2594591
// channel not read with timeout! 3048053
// channel not read with timeout! 3146052
// channel not read with timeout! 5358936
// channel not read with timeout! 5861733
// --- PASS: TestOriginal (5.96s)
func TestOriginal(t *testing.T) {
	for cnt := 0; cnt < count; cnt++ {
		x := make(chan bool, 1)
		x <- true
		select {
		case <-x:
		case <-time.After(time.Microsecond * 100):
			fmt.Printf("channel not read with timeout! %d\n", cnt)
			select {
			case <-x:
			default:
				panic("channel not read with default! ")
			}
		}
	}
}

// Let's do some profiling to understand how channels work and how fast they are.
//
// Histogram is a helper utility which helps us time repeatedly how much time is spent on operations.
// The buckets are powers of two (1,2,4,8,16...) and they are denoted in Nanosecond.
////////////////////////// HISTOGRAM BEGIN ////////////////
type Histogram struct {
	values    []int
	sum       int64
	n         int
	lastStart time.Time
}

// Start the timer.
func (h *Histogram) Start() {
	h.lastStart = time.Now()
}

// Stop timer, register value.
func (h *Histogram) Stop() {
	h.register(time.Since(h.lastStart).Nanoseconds())
}

// Prints histogram and average to stdout.
func (h *Histogram) Report() {
	p2 := 1
	for _, v := range h.values {
		if v != 0 {
			fmt.Println(comma(p2)+" ns", v)
		}
		p2 *= 2
	}
	fmt.Println("Average: ", h.sum/int64(h.n), " ns")
}

func (h *Histogram) register(v int64) {
	h.sum += v
	h.n++

	if h.values == nil {
		h.values = make([]int, 100) // 2^100 is big :)
	}
	bucket := int(math.Log2(float64(v)))
	if bucket < 0 {
		bucket = 0
	}
	h.values[bucket]++
}

// e.g. 1234567 -> "1,234,567"
func comma(v int) string {
	var comma string
	for i := v; i != 0; i /= 1000 {
		comma = "," + fmt.Sprintf("%03d", i%1000) + comma
	}
	comma = comma[1:]
	for comma[0] == '0' {
		comma = comma[1:]
	}
	return comma
}

////////////////////////// HISTOGRAM END ////////////////

// Before looking at channels, let's have a sense of how fast communicating through a slice would be.
// It's 6 ns
//
// $ go test -p=1 -v -run=Slice$ hello/selecttimeout_test.go
// === RUN   TestSlice
// Average:  6ns
func TestSlice(t *testing.T) {
	buf := []bool{}
	x := &buf
	start := time.Now()
	for i := 0; i < count; i++ {
		*x = append(*x, true)
		if len(*x) != 0 {
			*x = (*x)[1:]
		} else {
			t.Errorf("This should not happen")
		}
	}
	fmt.Println("Average: ", time.Since(start)/count)
}

// Let's see how the Histogram utility works.
// Notes:
//   - The histogram utility already has an overhead.
//       Now the average is 33ns instead fo 6!
//   - Has very long tail, with many above 1Microsecond. Some are even >100Microsecond
//
// $ go test -p=1 -v -run=SliceHist$ hello/selecttimeout_test.go
// === RUN   TestSliceHist
// 16, ns 8240201
// 32, ns 1067998
// 64, ns 682945
// 128, ns 6765
// 256, ns 708
// 512, ns 455
// 1,024, ns 550
// 2,048, ns 183
// 4,096, ns 71
// 8,192, ns 79
// 16,384, ns 29
// 32,768, ns 7
// 65,536, ns 5
// 131,072, ns 3
// 524,288, ns 1
// Average:  33  ns
func TestSliceHist(t *testing.T) {
	buf := []bool{}
	x := &buf
	hist := Histogram{}
	for i := 0; i < count; i++ {
		hist.Start()
		*x = append(*x, true)
		if len(*x) != 0 {
			*x = (*x)[1:]
		} else {
			t.Errorf("This should not happen")
		}
		hist.Stop()
	}
	hist.Report()
}

// Writing and reading a buffered channel with one thread.
//   - 74ns average time.
//   - This has a long tail too. Many is above 1 Microsec.
//       Interestignly the tail is smaller then for the slice. (more GC for slice?)
//
// $ go test -p=1 -v -run=BufferedOneThreadHist$ hello/selecttimeout_test.go
// === RUN   TestBufferedOneThreadHist
// 64, ns 9970143
// 128, ns 27991
// 256, ns 975
// 512, ns 69
// 1,024, ns 280
// 2,048, ns 216
// 4,096, ns 89
// 8,192, ns 147
// 16,384, ns 81
// 32,768, ns 8
// 65,536, ns 1
// Average:  74  ns
func TestBufferedOneThreadHist(t *testing.T) {
	c := make(chan struct{}, 1)
	hist := Histogram{}
	for i := 0; i < count; i++ {
		hist.Start()
		c <- struct{}{}
		<-c
		hist.Stop()
	}
	hist.Report()
}

// If we actually have two threads communicating on a buffered channel, it's much slower.
// Average is 216ns per operation
// We have a much bigger tail then in the one thread case .
//
// $ go test -p=1 -v -run=Buffered2ThreadsHist$ hello/selecttimeout_test.go
// === RUN   TestBuffered2ThreadsHist
// 256, ns 9162646
// 512, ns 741675
// 1,024, ns 62892
// 2,048, ns 26477
// 4,096, ns 5587
// 8,192, ns 401
// 16,384, ns 153
// 32,768, ns 91
// 65,536, ns 56
// 131,072, ns 7
// 262,144, ns 12
// 524,288, ns 1
// 1,048,576, ns 2
// Average:  432  ns
// Had 2 channel operations per element, so all values should be divided by 2!
func TestBuffered2ThreadsHist(t *testing.T) {
	c := make(chan struct{}, 1)
	d := make(chan struct{}, 1)
	go func() {
		for range c {
			d <- struct{}{}
		}
	}()
	hist := Histogram{}
	for i := 0; i < count; i++ {
		hist.Start()
		c <- struct{}{}
		<-d
		hist.Stop()
	}
	close(c)
	hist.Report()
	fmt.Println("Had 2 channel operations per element, so all values should be divided by 2!")
}

// Interesting takeaway, which might be counterintuitive:
//
// It's slightly faster using unbuffered channels then buffered ones!
// Also the tail is smaller.
// Of course we can not test the one thread original example with unbuffered channel.
//
// $ go test -p=1 -v -run=Unbuffered2ThreadsHist$ hello/selecttimeout_test.go
// === RUN   TestUnbuffered2ThreadsHist
// 256, ns 9384536
// 512, ns 540009
// 1,024, ns 47646
// 2,048, ns 24548
// 4,096, ns 2787
// 8,192, ns 276
// 16,384, ns 115
// 32,768, ns 57
// 65,536, ns 15
// 131,072, ns 7
// 262,144, ns 2
// 524,288, ns 2
// Average:  394  ns
// Had 2 channel operations per element, so all values should be divided by 2!
func TestUnbuffered2ThreadsHist(t *testing.T) {
	c := make(chan struct{})
	d := make(chan struct{})
	go func() {
		for range c {
			d <- struct{}{}
		}
	}()
	hist := Histogram{}
	for i := 0; i < count; i++ {
		hist.Start()
		c <- struct{}{}
		<-d
		hist.Stop()
	}
	close(c)
	hist.Report()
	fmt.Println("Had 2 channel operations per element, so all values should be divided by 2!")
}

// Summary so far: channels can be quite slow, they have a long tail.
//
// Let's slowly build up to the original example, to see which parts of it makes it slower
// and the tail longer.
//
// TestFinal1 is again a buffered channel with one thread, quite fast, but already long tail.
//
// $ go test -p=1 -v -run=Final1$ hello/selecttimeout_test.go
// === RUN   TestFinal1
// 64, ns 9982740
// 128, ns 15893
// 256, ns 670
// 512, ns 106
// 1,024, ns 192
// 2,048, ns 189
// 4,096, ns 25
// 8,192, ns 70
// 16,384, ns 43
// 32,768, ns 69
// 65,536, ns 2
// 131,072, ns 1
// Average:  70  ns
func TestFinal1(t *testing.T) {
	c := make(chan struct{}, 1)
	hist := Histogram{}
	for i := 0; i < count; i++ {
		hist.Start()
		c <- struct{}{}
		<-c
		hist.Stop()
	}
	hist.Report()
}

// What if we use select with default instead of just reading.
// Things to note:
//   - This never goes to the default case.
//       It means that go will make sure that the process of checking the channel is actually finished
//       before going to default.
//   - Using select is not slower than reading with <- normally from the channel (Final1)
//
// $ go test -p=1 -v -run=Final2$ hello/selecttimeout_test.go
// === RUN   TestFinal2
// 64, ns 9978716
// 128, ns 19786
// 256, ns 763
// 512, ns 108
// 1,024, ns 173
// 2,048, ns 183
// 4,096, ns 25
// 8,192, ns 103
// 16,384, ns 62
// 32,768, ns 76
// 65,536, ns 5
// Average:  72  ns
func TestFinal2(t *testing.T) {
	c := make(chan struct{}, 1)
	hist := Histogram{}
	for i := 0; i < count; i++ {
		hist.Start()
		c <- struct{}{}
		select {
		case <-c:
			hist.Stop()
		default:
			panic("This should not happen")
		}
	}
	hist.Report()
}

// This is like the original example.
// Select has a second timeout channel. This becomes super slow
// with very long tail up to 500 Milisecond!!!
//
// $ go test -p=1 -v -run=Final3$ hello/selecttimeout_test.go
// === RUN   TestFinal3
// 128, ns 1091
// 256, ns 8789040
// 512, ns 991233
// 1,024, ns 159794
// 2,048, ns 26315
// 4,096, ns 26366
// 8,192, ns 3670
// 16,384, ns 1233
// 32,768, ns 635
// 65,536, ns 150
// 131,072, ns 372
// 262,144, ns 33
// 524,288, ns 36
// 1,048,576, ns 17
// 2,097,152, ns 9
// 4,194,304, ns 2
// 8,388,608, ns 2
// 268,435,456, ns 1
// 536,870,912, ns 1
// Average:  532  ns
func TestFinal3(t *testing.T) {
	c := make(chan struct{}, 1)
	hist := Histogram{}
	for i := 0; i < count; i++ {
		hist.Start()
		c <- struct{}{}
		select {
		case <-c:
			hist.Stop()
		case <-time.After(time.Second):
			hist.Stop()
			log.Println("timeout")
		}
	}
	hist.Report()
}

// Actually Final3 was adding two new things to the timed region:
// - another branch for the select
// - it called time.After to have a channel.
//
// It turns out that if we don't time the call for time.After then
// the second select branch doesn't make it so much slower.
//
// $ go test -p=1 -v -run=Final4$ hello/selecttimeout_test.go
// === RUN   TestFinal4
// 64, ns 590255
// 128, ns 9256571
// 256, ns 141939
// 512, ns 1707
// 1,024, ns 2242
// 2,048, ns 4745
// 4,096, ns 2077
// 8,192, ns 379
// 16,384, ns 51
// 32,768, ns 19
// 65,536, ns 12
// 131,072, ns 1
// 262,144, ns 2
// Average:  157  ns
func TestFinal4(t *testing.T) {
	c := make(chan struct{}, 1)
	hist := Histogram{}
	for i := 0; i < count; i++ {
		timeout := time.After(time.Second)
		hist.Start()
		c <- struct{}{}
		select {
		case <-c:
			hist.Stop()
		case <-timeout:
			hist.Stop()
			log.Println("timeout")
		}
	}
	hist.Report()
}

// I'm not sure why time.After is so slow. I was curious if it creates separate goroutines too.
// The answer is NO additional goroutines are created.
// Well, it's probably just creating the returned channel and registering stuff.
//
// $ go test -p=1 -v -run=Goroutines$ hello/selecttimeout_test.go
// === RUN   TestGoroutines
// Number of goroutines before Sleep:  2
// Number of goroutines during Sleep:  10002
// Number of goroutines after Sleep:  6
// Number of goroutines before Timeafter:  2
// Number of goroutines during Timeafter:  2
// Number of goroutines after Timeafter:  2
func TestGoroutines(t *testing.T) {
	N := 10000
	finished := make(chan struct{})
	fmt.Println("Number of goroutines before Sleep: ", runtime.NumGoroutine())
	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(time.Second)
			finished <- struct{}{}
		}()
	}
	fmt.Println("Number of goroutines during Sleep: ", runtime.NumGoroutine())
	for i := 0; i < N; i++ {
		<-finished
	}
	fmt.Println("Number of goroutines after Sleep: ", runtime.NumGoroutine())
	time.Sleep(time.Millisecond * 10)

	fmt.Println("Number of goroutines before Timeafter: ", runtime.NumGoroutine())
	timeouts := [](<-chan time.Time){} // read only channels
	for i := 0; i < N; i++ {
		timeouts = append(timeouts, time.After(time.Second))
	}
	fmt.Println("Number of goroutines during Timeafter: ", runtime.NumGoroutine())
	for _, c := range timeouts {
		<-c
	}
	fmt.Println("Number of goroutines after Timeafter: ", runtime.NumGoroutine())
}

// Let's run the original code with the histogram.
// Although the timeout is 100 Microsecond, the actual time spent here can be much longer.
//
// My interpretation:
// a) x<-true will start a process, but not necessarily will be finished by this goroutine
//           (some outside scheduling might be needed).
// b) select starts to evaluate if there is data on the channels.
//           (The lookup for both channels might need outside scheduling)
// c) we just started time.After() right before this
//           which has big impact on scheduling and maybe increases likelihood for GC.
// d) we had some scheduling delay in the "case <-x", so we don't even know if there is a value there but
//    time.After already woke up and returned a value.
// f) select won't wait for evaluating the "case <-x" and will go the second branch.
//
// Also note that if there is no second branch the select will wait to fully evaluate the "case <-x"
// before going to default as demonstrated in the second select.
//
// $ go test -p=1 -v -run=OriginalHistogram$ hello/selecttimeout_test.go
// === RUN   TestOriginalHistogram
// channel not read with timeout! 120050
// channel not read with timeout! 897570
// channel not read with timeout! 1093143
// channel not read with timeout! 1429958
// channel not read with timeout! 2095051
// channel not read with timeout! 3265936
// channel not read with timeout! 4848782
// channel not read with timeout! 5893018
// channel not read with timeout! 5974106
// channel not read with timeout! 6245406
// 128, ns 22873
// 256, ns 8974033
// 512, ns 476089
// 1,024, ns 166934
// 2,048, ns 157457
// 4,096, ns 176391
// 8,192, ns 21816
// 16,384, ns 2120
// 32,768, ns 635
// 65,536, ns 911
// 131,072, ns 353
// 262,144, ns 310
// 524,288, ns 64
// 1,048,576, ns 13
// 4,194,304, ns 1
// Average:  560  ns
func TestOriginalHistogram(t *testing.T) {
	h := Histogram{}
	for cnt := 0; cnt < count; cnt++ {
		x := make(chan bool, 1)
		h.Start()
		x <- true
		select {
		case <-x:
			h.Stop()
		case <-time.After(time.Microsecond * 100):
			fmt.Printf("channel not read with timeout! %d\n", cnt)
			select {
			case <-x:
				h.Stop()
			default:
				panic("channel not read with default! ")
			}
		}
	}
	h.Report()
}

// Your workaround is good.
//
// My proposed solution is similar to yours.
// It will be better in case you expect that often there is value right away on the channel,
// because in that case it won't call time.After which could mess up scheduling.
func ReceiveWithTimeout(c chan bool, wait time.Duration) *bool {
	var b bool

	select {
	case b = <-c:
		return &b
	default:
	}

	select {
	case b = <-c:
		return &b
	case <-time.After(wait):
	}

	select {
	case b = <-c:
		return &b
	default:
	}

	return nil
}

// $ go test -p=1 -v -run=TestFixedHistogram$ hello/selecttimeout_test.go
// === RUN   TestFixedHistogram
// 64, ns 9756359
// 128, ns 236718
// 256, ns 3809
// 512, ns 340
// 1,024, ns 1272
// 2,048, ns 332
// 4,096, ns 395
// 8,192, ns 218
// 16,384, ns 184
// 32,768, ns 164
// 65,536, ns 207
// 131,072, ns 2
// Average:  95  ns
func TestFixedHistogram(t *testing.T) {
	h := Histogram{}
	for cnt := 0; cnt < count; cnt++ {
		x := make(chan bool, 1)
		h.Start()
		x <- true
		b := ReceiveWithTimeout(x, time.Microsecond*100)
		h.Stop()
		if b == nil || *b == false {
			panic("this should never happen! ")
		}
	}
	h.Report()
}
