package LargestNumber

func FindLargestNumber(numbers []int) int {
	Largest := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > Largest {
			Largest = numbers[i]
		}
	}

	return Largest
}
