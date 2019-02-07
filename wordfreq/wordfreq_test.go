package wordfreq

import (
	"fmt"
)

func ExampleGetWordCount() {
	wf := WordFreq{}
	// a := wordfreq.wor
	wf.AddWords("This test is a good test, or is it a a good example?")
	fmt.Println(wf.GetWordCount("example"))
	fmt.Println(wf.GetWordCount("test"))

	wf.AddWords("It's an example!")
	fmt.Println(wf.GetWordCount("example"))

	// Output:
	// 1
	// 2
	// 2
}
