package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func conc(str []byte, a chan int, v chan int, s chan int) {
	var alphabets, vowels, spaces int
	for i := 0; i < len(str); i++ {
		if (str[i] <= 95 && str[i] > 65) || (str[i] <= 122 && str[i] > 97) {
			alphabets += 1
			if str[i] == 97 || str[i] == 101 || str[i] == 105 || str[i] == 111 || str[i] == 117 {
				vowels += 1
			}
		} else if str[i] == 32 {
			spaces += 1
		}
	}
	// fmt.Print("\nhhhhhh")
	v <- vowels
	a <- alphabets
	s <- spaces
}
func main() {
	start := float64(time.Now().UnixNano()) / 1e9

	file, err := os.Open(os.Args[1])
	noOfPortions, err := strconv.Atoi(os.Args[2])
	var words, vowels, alphabets, spaces int = 0, 0, 0, 0
	if err != nil {
		fmt.Println(err)
	}
	str, err := io.ReadAll(file)
	len := len(str)
	a := make(chan int)
	v := make(chan int)
	s := make(chan int)
	for i := 1; i <= noOfPortions; i++ {
		go conc(str[(i-1)/noOfPortions*len:(i/noOfPortions)*len], a, v, s)
	}

	for i := 1; i <= noOfPortions; i++ {
		vowels += <-v
		alphabets += <-a
		spaces += <-s
	}

	words = spaces + 1
	end := float64(time.Now().UnixNano()) / 1e9

	fmt.Print("words: ", words, "\nVowels: ", vowels, "\nAlphabets: ", alphabets, "\n spaces: ", spaces, "\n Time taken: ", end-start)
	fmt.Println()
}
