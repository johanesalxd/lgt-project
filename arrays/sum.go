package arrays

func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	sums := make([]int, length)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
