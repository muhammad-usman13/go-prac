package main

import "fmt"

func mergeandsort(x, y []int) []int {
	merged := make([]int, len(x)+len(y))
	i, j := 0, 0
	for k := 0; k < len(merged); k++ {
		
		if i > len(x)-1 {
			merged[k] = y[j]
			j++
		} else if j > len(y)-1 {
			merged[k] = x[i]
			i++
		} else if x[i] < y[j] {
			merged[k] = x[i]
			i++
		} else {
			merged[k] = y[j]
			j++
		}
	}
	return merged
}	

func main() {
	x := []int{1, 2,2,2, 3}
	y := []int{1,2,2,3,3, 5, 6}
	fmt.Println(mergeandsort(x, y))
}