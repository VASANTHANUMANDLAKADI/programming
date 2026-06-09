package binarysearch

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{7, 8, 24, 32, 34, 66, 76, 98}
	target := 76

	assert.Equal(t, true, binarySearch(arr, target))
}

func TestBinarySearchNotFound(t *testing.T) {
	arr := []int{7, 8, 24, 32, 34, 66, 76, 98}

	target := 100

	assert.Equal(t, false, binarySearch(arr, target))
}