package main

func Fibonacci(n int)[]int{
	var first, second, tempresult int
	 result := []int{0,1}
		first = 0
		second = 1
		tempresult = first+second

	for n >= tempresult {
		tempresult = first +second
		if n >= tempresult {
			result= append(result,tempresult)
		}
		first = second
		second = tempresult
	} 
	return result
}