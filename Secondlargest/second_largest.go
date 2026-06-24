package secondlargest

func secondLargest(nums []int)int {

	if len(nums) < 2{
		return -1
	}

	largest := nums[0]					// Assigning 1st element in array as largest
	secondLargest := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {		// comparing the 1st elemtent in array with another elements in array
			secondLargest = largest
			largest = nums[i]
		} else if nums[i] > secondLargest && nums[i] != largest {  // checking weather secondlargest is lesser than largest element
			secondLargest = nums[i]
		}
		
	}

	return secondLargest
}
