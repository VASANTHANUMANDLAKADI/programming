package RemoveDuplicate
import (
		"testing"
		"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T){
		array := []int{1,2,4,6,9,4,3,4,6,7,6,3}
		expected := []int{1,2,4,6,9,3,7}

		assert.Equal(t, expected, removeDuplicates(array))
}