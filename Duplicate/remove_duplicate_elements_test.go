package removeduplicate
import (
		"testing"
		"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T){
		array := []int{1,2,4,6,9,4,3,4,6,7,6,3}
		expected := []int{1,2,4,6,9,3,7}

		assert.Equal(t, expected, RemoveDuplicates(array))
}

func TestEmptyArray(t *testing.T){
		array := []int{}
		expected := []int{}

		assert.Equal(t, expected, RemoveDuplicates(array))
}

func AllDuplicate(t *testing.T){
		array := []int{6,6,6,6,6,6,6,6,6,6}
		expected := []int{6}

		assert.Equal(t, expected, RemoveDuplicates(array))
}

func TestNoDuplicate(t *testing.T){
		array := []int{1,2,3,4,5,6,7,8,9,0}
		expected := []int{1,2,3,4,5,6,7,8,9,0}

		assert.Equal(t, expected, RemoveDuplicates(array))
}