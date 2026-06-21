package bubblesort

func BubbleSort(nums []int)[]int{
	n := len(nums)

	for i:= 0; i < n-1; i++ {		// This loop will iterate as per length of array
		for j:= 0; j< n-1-i; j++{	
			if nums[j] > nums[j+1]{  // Comparesion of elements in array
				nums[j], nums[j+1] = nums[j+1], nums[j] // Sorting of elements in array 
			}
		}
	}
	return nums
}
