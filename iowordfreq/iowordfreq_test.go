package iowordfreq

import "fmt"

func ExampleIoWordFreq() {
	iowf := IoWordFreq{}
	iowf.AddWords("hello friend")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))
	// Output:
	// 1
	// 2
	// 3
}
