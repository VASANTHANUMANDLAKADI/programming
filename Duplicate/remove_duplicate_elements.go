package removeduplicate

func RemoveDuplicates(nums []int)[]int {
	check := make(map[int]bool)
	result := []int{}

	for _, value := range nums {
		if !check[value] {
		    check[value] = true
		    result = append(result, value)
		}
	}
	return result
}


