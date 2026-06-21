package secondlargest
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondLargest(t*testing.T){
	nums:= []int{79,98,48,56,89}

	assert.Equal(t, 89, SecondLargest(nums))
}

func TestNegativeElementsArray(t*testing.T) {
		nums:= []int{-23,-56,56,87,-98}

		assert.Equal(t, 56, SecondLargest(nums))
}

func TestSingleElementArray(t*testing.T) {
		nums:= []int{24}

		assert.Equal(t, 1, SecondLargest(nums))
}

func TestDuplicateLargest(t*testing.T) {
		nums:= []int{5,5,5,5}

		assert.Equal(t, 1, SecondLargest(nums))
}

func TestEmptyArray(t*testing.T) {
		nums:= []int{}

		assert.Equal(t, 1, SecondLargest(nums))
}
