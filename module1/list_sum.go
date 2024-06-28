package module1

func ListSum(list []int) int {
	sum := 0

	for _, num := range list {
		sum += num
	}

	return sum
}

func ListSumRecursive(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + ListSumRecursive(list[1:])
}