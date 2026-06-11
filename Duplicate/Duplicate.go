package RemoveDuplicate

func removeDuplicates(array []int)[]int {
	check := make(map[int]bool)
	result := []int{}

	for _, value := range array {
		if !check[value] {
		    check[value] = true
		    result = append(result, value)
		}
	}
	return result
}


