package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, array := range numbersToSum {
		sums = append(sums, Sum(array))
	}

	return sums

}
