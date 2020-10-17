package prof

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func ExampleWordFreq() {
	iowf := WordFreq{}

	iowf.AddWords("hello friend")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprint(&iowf, "and hello again again")
	fmt.Println(iowf.GetWordCount("again"))

	fmt.Fprintln(&iowf, "cipő")
	fmt.Println(iowf.GetWordCount("cipő"))

	fmt.Fprintln(&iowf, "你，我")
	fmt.Println(iowf.GetWordCount("你"))
	fmt.Println(iowf.GetWordCount("我"))

	fmt.Fprintln(&iowf, "one，two")
	fmt.Println(iowf.GetWordCount("two"))

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 1
	// 1
	// 1
}

func ExampleFastWF() {
	iowf := FastWF{}
	iowf.AddWords("hello friend")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprint(&iowf, "and hello again again")
	fmt.Println(iowf.GetWordCount("again"))

	fmt.Fprintln(&iowf, "cipő")
	fmt.Println(iowf.GetWordCount("cipő"))

	fmt.Fprintln(&iowf, "你，我")
	fmt.Println(iowf.GetWordCount("你"))
	fmt.Println(iowf.GetWordCount("我"))

	fmt.Fprintln(&iowf, "one，two")
	fmt.Println(iowf.GetWordCount("two"))

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 1
	// 1
	// 1
}

func TestWithFile(t *testing.T) {
	buf, err := ioutil.ReadFile("../testdata/romeo_juliet.txt")
	if err != nil {
		t.Errorf("Couldn't open testdata file: %v", err)
		return
	}
	s := string(buf)
	iowf := WordFreq{}
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
		iowf := WordFreq{}
		iowf.Write(buf)
		if iowf.GetWordCount("Romeo") < 100 {
			b.Errorf(
				"Not enough words found Romeo: %d",
				iowf.GetWordCount("Romeo"))
		}
	}
}

func BenchmarkRomeoJulietFast(b *testing.B) {
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
