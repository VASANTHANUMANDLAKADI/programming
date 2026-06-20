Binary Search Overview :
Binary Search is an efficient searching algorithm that works on a sorted array. It repeatedly divides the search space in half until the target element is found or the search space becomes empty.

Execution :
1. Calculate the middle element
    mid := (left + right) / 2
2. Compare the middle element with the target
    If equal, return true
    If target is smaller, search the left half
    If target is larger, search the right half
3. Continue until the target is found or the search range becomes empty.

Unit Tests :

TestBinarySearch:
Verifies that the function returns true when the target element exists in the sorted array.

TestBinarySearchNotFound:
Verifies that the function returns false when the target element does not exist in the sorted array.