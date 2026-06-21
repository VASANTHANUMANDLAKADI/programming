package main
import "fmt"

func mergesortedarrays(arr1,arr2 []int)[]int{
	result:= []int{}
	i:= 0
	j:= 0

	for i < len(arr1) && j < len(arr2){
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
		for i < len(arr1) {
			result = append(result, arr1[i])
			i++
		}

		for j < len(arr2){
			result = append(result, arr2[j])
			j++
		}
		return result
	}

func main(){
	arr1 := []int{32,65,7,23,32,67,89}
	arr2 := []int{54,56,22,87,98,35}

	fmt.Println(mergesortedarrays(arr1,arr2))
}