package secondlargest

func SecondLargest(num []int) int {
	largest := num[0]					// Assigning 1st element in array as largest
	secondlargest := num[0]

	for i := 1; i < len(num); i++ {
		if num[i] > largest {		// comparing the 1st elemtent in array with another elements in array
			secondlargest = largest
			largest = num[i]
		} else if  num[i] > secondlargest && num[i] != largest {  // checking weather secondlargest is lesser than largest element
			secondlargest = num[i]
		}
	}

	return secondlargest
}
