package module2

// Slice Tricks: https://go.dev/wiki/SliceTricks
func InsertionSortInt(list []int) {
	for i := 1; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[i] <= list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
