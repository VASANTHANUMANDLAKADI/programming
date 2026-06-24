package secondlargest
import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestSecondLargest(t*testing.T){
	num := []int{79,98,48,56,89}

	assert.Equal(t, 89, secondLargest(num))
}

func TestSecondLargestNegative(t*testing.T) {
		num := []int{-23,-56,56,87,-98}

		assert.Equal(t, 56, secondLargest(num))
}

func TestSingleElementArray(t*testing.T) {
		num := []int{24}

		assert.Equal(t, -1, secondLargest(num))
}

func TestDuplicateLargest(t*testing.T) {
		num := []int{24,56,66,45,56}

		assert.Equal(t, 56, secondLargest(num))
}

func TestEmptyElementArray(t*testing.T) {
		num := []int{}

		assert.Equal(t, -1, secondLargest(num))
}
