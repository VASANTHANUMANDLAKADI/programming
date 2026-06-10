package main

func bubbleSort(array []int)[]int{
	n := len(array)

	for i:= 0; i < n-1; i++ {		// This loop will iterate as per length of array
		for j:= 0; j< n-1-i; j++{	
			if array[j] > array[j+1]{  // Comparesion of elements in array
				array[j], array[j+1] = array[j+1], array[j] // Sorting of elements in array 
			}
		}
	}
	return array
}
