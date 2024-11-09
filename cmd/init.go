/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/albertapi/AlbertApiCLI/sets"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new testing set",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.MkdirAll(Path, 0755)

		set := sets.Set{
			Name:      "",
			Config:    "",
			Data:      "",
			DataLabel: "",
		}

		sets.SaveSet(Path, set)

		os.MkdirAll(Path + "/FindReplaces", 0755)
		os.MkdirAll(Path + "/Requests", 0755)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCmd.Flags().StringVarP(&Path, "Path", "P", "", "The path to create the new set (Required)")
	initCmd.MarkFlagRequired("Path")
}
