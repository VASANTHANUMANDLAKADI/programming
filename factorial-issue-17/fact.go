package main

func Factorial(n int)int{

	factorial := 1

	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}

	return factorial 
}
