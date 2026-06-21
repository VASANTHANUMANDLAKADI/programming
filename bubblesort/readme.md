BubbleSort :
Bubble Sort is a simple sorting algorithm that repeatedly compares adjacent elements and swaps them if they are in the wrong order. The process continues until the entire array is sorted.
If the current element is greater than the next element, they are swapped.

The Bubble Sort function:

Iterates through the array multiple times.
Compares adjacent elements.
Swaps elements when they are out of order.
Returns the sorted array in ascending order.

Unit Tests:

TestBubbleSort :
Verifies that an unsorted array is correctly sorted in ascending order.
The test compares the actual sorted result with the expected output using assertions.
Example:

input    := []int{66,44,75,2,29,20,24}
expected := []int{2,20,24,29,44,66,75}