package module2

func BubbleSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list) - i - 1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
