package bubblesort

import (
		"testing"
		"github.com/stretchr/testify/assert"
	)
func TestBubbleSort(t *testing.T){

	nums := []int {66,44,75,2,29,20,24}
	expected := []int {2,20,24,29,44,66,75}

	assert.Equal(t, expected, BubbleSort(nums))
	
}

func TestAlreadysorted(t *testing.T){

	nums := []int {1,2,3,4,5,6}
	expected := []int {1,2,3,4,5,6}

	assert.Equal(t, expected, BubbleSort(nums))
	
}

func TestReverseSorted(t *testing.T){

	nums := []int {75,66,44,29,24,20,2}
	expected := []int {2,20,24,29,44,66,75}

	assert.Equal(t, expected, BubbleSort(nums))
	
}

func TestDuplicateElements(t *testing.T){

	nums := []int {66,24,75,2,29,66,24}
	expected := []int {2,24,24,29,66,66,75}

	assert.Equal(t, expected, BubbleSort(nums))
	
}
