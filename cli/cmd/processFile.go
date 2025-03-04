/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// processFileCmd represents the processFile command
var processFileCmd = &cobra.Command{
	Use:   "processFile",
	Short: "Process a file",
	Long: `Process a file in chunks
	For Example :
	cli processFile -f ~/code/go prac/FileProcessing/large_text_2.txt -s 1`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		chunkSize, _ := cmd.Flags().GetInt("chunkSize")
		mainprocess(file, chunkSize)
	},
}

func init() {
	rootCmd.AddCommand(processFileCmd)

	// Here you will define your flags and configuration settings.
	processFileCmd.Flags().StringP("file", "f", "~/code/go prac/FileProcessing/large_text_2.txt", "file path")
	processFileCmd.Flags().IntP("chunkSize", "s", 1, "chunk size")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
	v <- vowels
	a <- alphabets
	s <- spaces
}
func mainprocess(fileToProcess string, noOfPortions int) {
	start := float64(time.Now().UnixNano()) / 1e9
	file, err := os.Open(fileToProcess)
	var words, vowels, alphabets, spaces int = 0, 0, 0, 0
	if err != nil {
		fmt.Println(err)
	}
	str, _ := io.ReadAll(file)
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
