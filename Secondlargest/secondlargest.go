package main

func secondLargest(num []int) int {
	largest := num[0]
	secondLargest := num[0]

	for i := 1; i < len(num); i++ {
		if num[i] > largest {
			secondLargest = largest
			largest = num[i]
		} else if num[i] > secondLargest && num[i] != largest {
			secondLargest = num[i]
		}
	}

	return secondLargest
}
