package arrays

// Sum returns sum of all the numbers in an array
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
