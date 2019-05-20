package wordfreq

import "fmt"

func ExampleGetWordCount() {
	var wf WordFreq
	fmt.Println("zero:", wf.GetWordCount("zero"))

	wf.AddWords("Is this a test or just an example?")
	wf.AddWords("It's an example.")

	fmt.Println("test:", wf.GetWordCount("test"))
	fmt.Println("example:", wf.GetWordCount("example"))

	wf.AddWords("Looks like a test to me.")
	wf.AddWords("Test? It's a TEST!")

	fmt.Println("test:", wf.GetWordCount("test"))
	fmt.Println("example:", wf.GetWordCount("example"))

	wf.AddWords("This is last")
	fmt.Println("last:", wf.GetWordCount("last"))

	// Output:
	// zero: 0
	// test: 1
	// example: 2
	// test: 4
	// example: 2
	// last: 1
}
