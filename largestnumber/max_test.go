package main
import "testing"

func TestFindMax(t *testing.T) {

	tests := []struct {
		input    []int
		expected int
	}{{[]int{3, 7, 2, 9}, 9},{[]int{10, 5, 1}, 10},{[]int{2, 4, 8}, 8},}

	for _, test := range tests {

		result := FindMax(test.input)

		if result != test.expected {
			t.Errorf("Expected %d but got %d",test.expected,result,)
		}
	}
}