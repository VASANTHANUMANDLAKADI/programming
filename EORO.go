package main
import ("fmt")

func main() {

	var a int

	for{
  
    fmt.Println("Enter a number:")

	_,err := fmt.Scan(&a)

	if err != nil {
		fmt.Println("Invalid input")
		var dummy string
		fmt.Scan(&dummy)
		continue
	}
	break
	}
	
    if a%2 == 0 {
    
    fmt.Println("Number is even")

  } else { 
  	fmt.Println("Number is odd")
  }
}
