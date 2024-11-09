
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/albertapi/AlbertApiCLI/findreplaces"
	"github.com/albertapi/AlbertApiCLI/sets"
)

// addFindReplaceCmd represents the addFindReplace command
var addFindReplaceCmd = &cobra.Command{
	Use:   "addFindReplace",
	Short: "Add a new find and replace",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := Name + ".json"

		findReplace := findreplaces.FindReplace{
			FileName: fileName,
			Name: Name,
		}

		if findreplaces.FindReplaceExists(Path, fileName) {
			fmt.Println("Find Replace exists. Overwrite (y/n)?")
			var overwriteString string
			fmt.Scanln(&overwriteString)
			if overwriteString != "y" {
				return
			}
		}

		findreplaces.SaveFindReplace(Path, fileName, findReplace)

		set := sets.LoadSet(Path)

		if set.FindReplaces == nil {
			set.FindReplaces = make([]string, 0)
		}

		set.FindReplaces = append(set.FindReplaces, fileName)

		sets.SaveSet(Path, set)
	},
}

func init() {
	rootCmd.AddCommand(addFindReplaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addFindReplaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addFindReplaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addFindReplaceCmd.Flags().StringVarP(&Path, "Path", "P", "", "The path to the set")
        addFindReplaceCmd.MarkFlagRequired("Path")

        addFindReplaceCmd.Flags().StringVarP(&Name, "Name", "N", "", "The name of the request")
        addFindReplaceCmd.MarkFlagRequired("Name")
}
