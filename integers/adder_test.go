package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	sum := Add(2, 1)
	expexcted := 3

	if sum != expexcted {
		t.Errorf("expected %d but got %d", expexcted, sum)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
