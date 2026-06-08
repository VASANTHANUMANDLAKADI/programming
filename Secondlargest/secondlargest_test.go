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