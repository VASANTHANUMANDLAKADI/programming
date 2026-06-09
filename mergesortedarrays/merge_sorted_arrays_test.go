package merge
import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestMergesortedarrays(t *testing.T){
	a := []int{14,22,24,56,99}
	b := []int{16,36,67,84,100}

	expected:= []int{14,16,22,24,36,56,67,84,99,100}

	assert.Equal(t, expected, mergesortedarrays(a,b))
}

