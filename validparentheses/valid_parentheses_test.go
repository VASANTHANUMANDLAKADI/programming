package parentheses
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVaildParentheses(t* testing.T){
	a:= "[{()}]"

	assert.Equal(t, true, isBalanced(a))
}

func TestEmptyString(t* testing.T){
	a:= ""

	assert.Equal(t, true, isBalanced(a))
}

func TestIncorrectOrder(t* testing.T){
	a:= "[({)]}"

	assert.Equal(t, false, isBalanced(a))
}

func TestMissingParentheses(t* testing.T){
	a:= "[{(]"

	assert.Equal(t, false, isBalanced(a))
}