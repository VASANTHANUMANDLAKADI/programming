package main
import "testing"

func TestSecondLargest(t*testing.T){
	num := []int{79,98,48,56,89}

	result := secondLargest(num)

	Expected := 89

	if result != 89 {
		t.Errorf("Expected %d,but got %d", Expected, result)
	}
}

func TestSecondLargestNegative(t*testing.T) {
		num := []int{-23,-56,56,87,-98}

		result := secondLargest(num)

		Expected := 56

	if result != 56 {
		t.Errorf("Expected %d, got %d", Expected, result)
	}
}

func TestSingleElementArray(t*testing.T) {
		num := []int{24}

		result := secondLargest(num)

		Expected := 24

	if result != 24 {
		t.Errorf("Expected %d, got %d", Expected, result)
	}
}

func TestDuplicateLargest(t*testing.T) {
		num := []int{24,56,66,45,66}

		result := secondLargest(num)

		Expected := 56

	if result != 56 {
		t.Errorf("Expected %d, got %d", Expected, result)
	}
}
