package main

func FindMax(numbers []int) int {
	max := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return max
}
