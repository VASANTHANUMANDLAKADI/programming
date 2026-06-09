package merge

func mergesortedarrays(a,b []int)[]int{
	result:= []int{}
	i:= 0
	j:= 0

	for i<len(a)&& j<len(b){
		if a[i]<b[j]{
			result = append(result,a[i])
			i++
		}else{
			result = append(result,b[j])
		}
	}
	for i<len(a){
		result = append(result,a[i])
		i++
	}
	for j<len(b){
		result = append(result,b[j])
		j++
	}
	return result
}