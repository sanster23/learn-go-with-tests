package arrays

// Sum returns sum of all the numbers in an array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll adds all the slices and returns the total of slices in a new slice
func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails adds all tails from given slices and returns a new slice
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
