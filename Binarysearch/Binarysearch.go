package binary

func binarySearch(array []int, target int)bool {
	left := 0
	right := len(array)-1

	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return true
		}

		if target < array[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}