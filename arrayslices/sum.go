package arrslice

func Sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(tailsToChase ...[]int) []int {
	var tails []int
	for _, num := range tailsToChase {
		tails = append(tails, num[len(num)-1])
	}

	return tails
}
