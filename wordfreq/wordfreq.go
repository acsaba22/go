package wordfreq

// Make:
// type WordFreq struct {...}
//
// And make the test compile, then pass.
//
// To split a string into pieces you can iterate over:
// buf := bytes.Buffer{}
// for over string
//   if unicode.IsLetter(r)
//      buf.WriteRune(unicode.ToLower(r))
//   else
//      we have one token in buf.String
//      buf.Reset
