package wordfreq

import (
	"fmt"
)

func ExampleGetWordCount() {
	wf := WordFreq{}
	// a := wordfreq.wor
	wf.AddWords("Is this a test or just an example?")
	wf.AddWords("It's an example.")

	fmt.Println("test:", wf.GetWordCount("test"))
	fmt.Println("example:", wf.GetWordCount("example"))

	wf.AddWords("Looks like a test to me.")
	wf.AddWords("Test? It's a TEST!")

	fmt.Println("test:", wf.GetWordCount("test"))
	fmt.Println("example:", wf.GetWordCount("example"))

	// Output:
	// test: 1
	// example: 2
	// test: 4
	// example: 2
}
