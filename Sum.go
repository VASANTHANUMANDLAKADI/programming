package main
import "fmt"

func main() {
	var N int
	sum := 0

	for {
		fmt.Println("Enter a number")
		_,err := fmt.Scan(&N)

		if err != nil {
			fmt.Println("Invalid input")
			var dummy string
			fmt.Scan(&dummy)
			continue
		}
		break
	}

	for i := 1; i <= N; i++ {
			sum = sum + i
	}

	fmt.Println(sum)
}
