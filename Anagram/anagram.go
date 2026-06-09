package anagram

func Anagram(a string, b string)bool {
	if len(a) != len(b) {     //Checking lengths of 2 strings
		return false          //If the lengths are not equal testcase fails
	}

	count := make(map[rune]int)

	// Assigning vales to each character
	//Iteration through each character
	for _, ch := range a {
		count[ch]++
	}

	for _, ch := range b {      
		count[ch]--
	}

	for _, value := range count {        
		if value != 0 {
			return false
		}
	}
	return true
}