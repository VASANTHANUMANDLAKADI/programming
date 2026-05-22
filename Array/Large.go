package main
import "fmt"

func main() {

	numbers := [5]int{32,421,345,546,123}

	max := numbers[0]

	for i := 0; i < len(numbers); i++{
		
		if numbers[i] > max{
		max = numbers[i]
		}
	}
	fmt.Println("Largest number in array is:",max)
}
