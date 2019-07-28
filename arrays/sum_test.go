package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 5}, []int{2, 6})
	want := []int{6, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	}

	t.Run("make the sum of all slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 5})
		want := []int{5, 8}
		checkSum(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 5})
		want := []int{0, 8}
		checkSum(t, got, want)
	})
}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 5, 3, 10}))
	// Output: 21
}

func ExampleSumAll() {
	fmt.Println(SumAll([]int{1, 2}, []int{3, 4}))
	// Output: [3 7]
}

func ExampleSumAllTails() {
	fmt.Println(SumAllTails([]int{1, 2, 5}, []int{2, 3, 10}))
	fmt.Println(SumAllTails([]int{}, []int{2, 3, 11}))
	// Output:
	// [7 13]
	// [0 14]
}
