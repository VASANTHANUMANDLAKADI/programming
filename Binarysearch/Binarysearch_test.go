package binary

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{7, 8, 24, 32, 34, 66, 76, 98}

	target := 34

	result := binarySearch(arr, target)

	if result != true {
		t.Errorf("Expected true, got false")
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	arr := []int{7, 8, 24, 32, 34, 66, 76, 98}

	target := 100

	result := binarySearch(arr, target)

	if result != false {
		t.Errorf("Expected false, got true")
	}
}