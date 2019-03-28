package prof

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/acsaba22/go/iowordfreq"
)

func ExampleIoWordFreq() {
	iowf := iowordfreq.IoWordFreq{}

	iowf.AddWords("hello friend")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again again")
	fmt.Println(iowf.GetWordCount("hello"))
	// Output:
	// 1
	// 2
	// 3
}

func ExampleFastWF() {
	iowf := FastWF{}
	iowf.AddWords("hello friend")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again again")
	fmt.Println(iowf.GetWordCount("hello"))
	// Output:
	// 1
	// 2
	// 3
}

func TestWithFile(t *testing.T) {
	buf, err := ioutil.ReadFile("../testdata/romeo_juliet.txt")
	if err != nil {
		t.Errorf("Couldn't open testdata file: %v", err)
		return
	}
	s := string(buf)
	iowf := iowordfreq.IoWordFreq{}
	fmt.Fprint(&iowf, s)
	if iowf.GetWordCount("Romeo") < 100 || iowf.GetWordCount("Juliet") < 100 {
		t.Errorf(
			"Not enough words found Romeo: %d ; Juliet: %d",
			iowf.GetWordCount("Romeo"),
			iowf.GetWordCount("Juliet"))
	}
}

func TestWithFileFast(t *testing.T) {
	buf, err := ioutil.ReadFile("../testdata/romeo_juliet.txt")
	if err != nil {
		t.Errorf("Couldn't open testdata file: %v", err)
		return
	}
	s := string(buf)
	iowf := FastWF{}
	fmt.Fprint(&iowf, s)
	if iowf.GetWordCount("Romeo") < 100 || iowf.GetWordCount("Juliet") < 100 {
		t.Errorf(
			"Not enough words found Romeo: %d ; Juliet: %d",
			iowf.GetWordCount("Romeo"),
			iowf.GetWordCount("Juliet"))
	}
}

func BenchmarkRomeoJuliet(b *testing.B) {
	fmt.Println("Benchmark start")
	buf, err := ioutil.ReadFile("../testdata/romeo_juliet.txt")
	if err != nil {
		b.Errorf("Couldn't open testdata file: %v", err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iowf := iowordfreq.IoWordFreq{}
		iowf.Write(buf)
		if iowf.GetWordCount("Romeo") < 100 {
			b.Errorf(
				"Not enough words found Romeo: %d",
				iowf.GetWordCount("Romeo"))
		}
	}
}

func BenchmarkRomeoJulieFast(b *testing.B) {
	fmt.Println("Benchmark start")
	buf, err := ioutil.ReadFile("../testdata/romeo_juliet.txt")
	if err != nil {
		b.Errorf("Couldn't open testdata file: %v", err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iowf := FastWF{}
		iowf.Write(buf)
		if iowf.GetWordCount("Romeo") < 100 {
			b.Errorf(
				"Not enough words found Romeo: %d",
				iowf.GetWordCount("Romeo"))
		}
	}
}
