package main
import (
		"testing"
		"reflect"
)

func TestRemoveDuplicate(t *testing.T){
		array := []int{1,2,4,6,9,4,3,4,6,7,6,3}
		expected := []int{1,2,4,6,9,3,7}

		result := removeDuplicates(array)

		if !reflect.DeepEqual(result,expected){
			t.Errorf("Expected %v but got %v", expected, result)
		}
}