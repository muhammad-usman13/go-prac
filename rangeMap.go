package main

import "fmt"

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10}
	fmt.Print(m)
	for key, value := range m {
		println(key, value)
	}
}
