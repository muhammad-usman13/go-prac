package main

import "fmt"
func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
