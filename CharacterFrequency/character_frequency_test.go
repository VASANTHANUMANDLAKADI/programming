package characterfrequency
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFrequencyCount(t *testing.T){
	expected:= map[rune]int{
    'b': 1,
    'a': 3,
    'n': 2,
} 
	assert.Equal(t,expected,FrequencyCount("banana"))
}

func FrequencyCountUpperAndLowerCase(t *testing.T){
	expected:= map[rune]int{
	'b': 1,
    'a': 3,
    'n': 2,
	}
	assert.Equal(t,expected,FrequencyCount("BAnaNa"))
}

func TestFrequencyCountSpecialCharacters(t *testing.T) {

	expected := map[rune]int{
		'@': 2,
		'#': 1,
		'$': 1,
	}

	assert.Equal(t, expected, FrequencyCount("@@#$"))
}

func TestFrequencyCountEmptyString(t *testing.T) {

	expected := map[rune]int{}

	assert.Equal(t, expected, FrequencyCount(""))
}