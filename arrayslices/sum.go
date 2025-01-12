package arrslice

func Sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}

func SumAll(a, b []int) []int {
	var outslice []int
	outslice = append(outslice, Sum(a))
	outslice = append(outslice, Sum(b))
	return outslice
}
