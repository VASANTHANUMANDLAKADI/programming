package parentheses
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVaildParentheses(t* testing.T){
	a:= "[{()}]"

	assert.Equal(t, true, isBalanced2(a))
}

func TestEmptyString(t* testing.T){
	a:= ""

	assert.Equal(t, true, isBalanced2(a))
}

func TestIncorrectOrder(t* testing.T){
	a:= "[({)]}"

	assert.Equal(t, false, isBalanced2(a))
}

func TestMissingParentheses(t* testing.T){
	a:= "[{(]"

	assert.Equal(t, false, isBalanced2(a))
}

func TestMissingParentheses_single(t* testing.T){
	a:= "[{}]"
	
	// split 
	// s1 = "[{"
	// s2 = "}]"
	//counter = len(s2), counter = conter -1
	//pairs [<===>], {<==>}

	//loop s1 
	//1st -> [
	   //check it's pair, its is ]
	   //get last value from s2, s2[1] == ]
	   //if s1[1st] != s2[last] return false
	//2nd -> {
	  //check it's pair, its is }
	   //get last second value from s2, s2[1] == ]
	   //if s1[1st] != s2[last] return false
	  // counter = counter -1
	assert.Equal(t, true, isBalanced2(a))
}