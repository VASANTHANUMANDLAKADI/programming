package main
import "testing"

func TestIsPrime(t *testing.T){
	tests := []struct {
		input int
		 expected bool
	}	{
			{2, true},
			{3, true},
			{4, false},
			{5, true},
			{6, false},
			{7, true},
			{8, false},
			{11, true},
			{1, false},
			{0, false},
			{-5, false},
		}

		for _, test := range tests {
			result := IsPrime(test.input)

			if result != test.expected {
				t.Errorf("IsPrime(%d) = %v; expected %v",
				test.input, result, test.expected)
			}
		}
}