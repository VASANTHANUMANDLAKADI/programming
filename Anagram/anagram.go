package anagram

func Anagram(a string, b string)bool {
	if len(a) != len(b) {			// checking weather both strings have equal lengths
		return false
	}

	count := make(map[rune]int)

	for _, ch := range a {   // Iteration through each character in strings
		count[ch]++			// Assigning Values to each character
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