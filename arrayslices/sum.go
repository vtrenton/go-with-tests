package arrslice

func Sum(nums [5]int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}
