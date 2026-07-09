package parentheses

//s = []
func isBalanced(str string)bool{
	stack:= []string{}
	
	// if len(s) %2 != 0 {
	// 	return false
	// }
	pairs:= map[string]string{
		")" : "(",
		"}" : "{",
		"]" : "[",
	}
	for _, ch := range str {
		s := string(ch)
		switch s {
		case "(", "{", "[":
			stack = append(stack, s) //stack = [

		case ")", "}", "]":
			if len(stack) == 0{
				return false
			}
			 top:= stack[len(stack)-1] //top = 
			 stack = stack[:len(stack)-1] //empty

			 if top != pairs[s]{
				return false
			 }
		}
	}
	return len(stack) == 0
}

func isBalanced2(s string) bool {

	pairs:= map[string]string{
		")" : "(",
		"}" : "{",
		"]" : "[",
	}

	//mid := len(s) /2
	for i := 0; i < len(s) /2; i++ {
		
		lastIndex := (len(s) -i) -1

		lastChar := string(s[lastIndex])
		firstChar := string(s[i])
		if firstChar != pairs[lastChar] {
			return false
		}
	}
   
	return true
}


