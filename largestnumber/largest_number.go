package largestnumber

func FindLargestNumber(numbers []int)int {
	largest := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
	}

	return largest
}
