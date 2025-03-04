/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hexCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Generate Hex numbers",
	Long: `provided length it generates hex numbers
	For Example :
	cli hex -l 4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hex called")
	},
}

func init() {
	rootCmd.AddCommand(hexCmd)

	// Here you will define your flags and configuration settings.
	hexCmd.Flags().IntP("length", "l", 4, "Length of hex")
	
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
