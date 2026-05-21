package main

import "fmt"

func main() {
	var n int

	for{

	fmt.Println("Enter the number to print the multiplication table: ")

	_,err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Invalid input")
		var dummy string
		fmt.Scan(&dummy)
		continue
		}
	break
	}
	fmt.Printf("Multiplication Table for %d:\n", n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}