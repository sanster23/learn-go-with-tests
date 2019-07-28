package arrays

import (
	"fmt"
	"testing"
)

func TestSumSlices(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 5, 3, 10}))
	// Output: 21
}
