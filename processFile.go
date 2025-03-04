package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	start := float64(time.Now().UnixNano()) / 1e9

	file, err := os.Open(os.Args[1])
	var words, vowels, alphabets, spaces int
	if err != nil {
		fmt.Println(err)
	}
	str, err := io.ReadAll(file)
	c := len(str)
	for i := 0; i < c; i++ {
		if (str[i] <= 95 && str[i] > 65) || (str[i] <= 122 && str[i] > 97) {
			alphabets += 1
			if str[i] == 97 || str[i] == 101 || str[i] == 105 || str[i] == 111 || str[i] == 117 {
				vowels += 1
			}
		} else if str[i] == 32 {
			spaces += 1
		}
	}
	words = spaces + 1
	end := float64(time.Now().UnixNano()) / 1e9

	fmt.Print("words: ", words, "\nVowels: ", vowels, "\nAlphabets: ", alphabets, "\n spaces: ", spaces, "\n Time taken: ", end-start)

	fmt.Println()
}
