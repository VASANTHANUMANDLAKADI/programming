package main
import "testing"

func TestFactorial(t *testing.T) {
	
	result := Factorial(5)
	
	if result != 120 {
		t.Errorf("Expected 120 but got %d",result)
	}

}