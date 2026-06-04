package main

import (
		"testing"
		"reflect"
)
func TestBubbleSort(t *testing.T){

	array := []int {66,44,75,2,29,20,24}
	expected := []int {2,20,24,29,44,66,75}

	array = bubbleSort(array) 

	if !reflect.DeepEqual(array,expected) {
		t.Errorf("Expected %v but got %v", expected,array)
	}
}
