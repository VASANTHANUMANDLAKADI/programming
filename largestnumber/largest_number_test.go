package largestnumber
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindLargestNumber(t *testing.T) {

	array := []int{2,4,6,3,7,24}

	assert.Equal(t, 24, FindLargestNumber(array))
}