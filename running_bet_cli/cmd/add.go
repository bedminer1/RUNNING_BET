/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/bedminer1/running_bet/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a record",

	RunE: func(cmd *cobra.Command, args []string) error {
		myScore := float32(viper.GetFloat64("myScore"))
		herScore := float32(viper.GetFloat64("herScore"))

		return addAction(os.Stdout, "record", myScore, herScore)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Float64P("myScore", "m", 0.00, "How much I ran that week")
	addCmd.Flags().Float64P("herScore", "y", 0.00, "How much Yoona ran that week")

	viper.BindPFlag("myScore", addCmd.Flags().Lookup("myScore"))
	viper.BindPFlag("herScore", addCmd.Flags().Lookup("herScore"))
}

func addAction(out io.Writer, file string, myScore, herScore float32) error {
	db := &api.Records{}
	if err := db.Get("local_storage/", file, "json"); err != nil {
		return err
	}

	if err := db.Add(myScore, herScore, [][]float32{}); err != nil {
		return err
	}

	if err := db.Save("local_storage/", file); err != nil {
		return err
	}

	fmt.Fprintf(out, "Added item\n")
	return nil
}
