package main
import "fmt"

func main() {
	m := make([]int, 6)

	for i := 0; i < 5; i++ {
		m[i] = i
	}
	fmt.Println(len(m), cap(m), m)
	
	m = append(m, 5, 3,89,99,90)
	fmt.Println(len(m), cap(m), m)
	slicedArr := m[0:7]
	append(slicedArr, 10)
	fmt.Println(len(slicedArr), cap(slicedArr), slicedArr)
}