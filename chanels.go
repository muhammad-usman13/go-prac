package main

import "fmt"

func sum(nums []int, result chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	result <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := make(chan int)
	go sum(nums[:len(nums)/2], result)
	go sum(nums[len(nums)/2:], result)
	x, y := <-result, <-result
	fmt.Println(x, y, x+y)
}
