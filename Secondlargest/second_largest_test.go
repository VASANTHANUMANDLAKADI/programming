package secondlargest
import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestSecondLargest(t*testing.T){  // Tests a normal array with positive numbers.
	nums:= []int{79,98,48,56,89}

	assert.Equal(t, 89, secondLargest(nums))
}

func TestSecondLargestNegative(t*testing.T) {  // Tests an array containing both positive and negative numbers.
		nums:= []int{-23,-56,56,87,-98}

		assert.Equal(t, 56, secondLargest(nums))
}

func TestSingleElementArray(t*testing.T) {  // Tests an array with only one element.
		nums:= []int{24}

		assert.Equal(t, -1, secondLargest(nums))
}

func TestDuplicateLargest(t*testing.T) {  // Tests an array containing duplicate values.
		nums:= []int{24,56,66,45,56}

		assert.Equal(t, 56, secondLargest(nums))
}

func TestEmptyElementArray(t*testing.T) {  // Tests an empty array.
		nums:= []int{}

		assert.Equal(t, -1, secondLargest(nums))
}
