package main

import "fmt"

func main(){
	var n int
	rev := 0

	for {
	fmt.Println("Enter a number:")
	_,err := fmt.Scan(&n)

	if err != nil {

		fmt.Println("Invalid input")
		var dummy string
		fmt.Scan(&dummy)
		continue
	}
	break
	}

	for n > 0{
	digit := n%10
	rev = rev*10 + digit
	n = n/10
	}

	fmt.Println("Reversed number", rev)
	
}