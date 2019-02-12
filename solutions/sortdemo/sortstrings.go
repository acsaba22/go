package sortdemo

type SSlice []string

func (ss SSlice) Len() int {
	return len(ss)
}

func (ss SSlice) Less(i, j int) bool {
	return ss[i] < ss[j]
}
func (ss SSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
