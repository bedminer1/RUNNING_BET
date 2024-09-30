/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a record",
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Float32P("myScore", "m", 0.00, "How much I ran that week")
	addCmd.Flags().Float32P("herScore", "y", 0.00, "How much Yoona ran that week")
}
