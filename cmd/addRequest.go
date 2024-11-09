/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/albertapi/AlbertApiCLI/sets"
	"github.com/albertapi/AlbertApiCLI/requests"
)

// addRequestCmd represents the addRequest command
var addRequestCmd = &cobra.Command{
	Use:   "addRequest",
	Short: "Add a new request to the list",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := Name + ".json"

		request := requests.Request{
			FileName: fileName,
			Name: Name,
		}

		requests.SaveRequest(Path, fileName, request)

		set := sets.LoadSet(Path)

		if set.Requests == nil{
			set.Requests = make([]string, 0)
		}

		set.Requests = append(set.Requests, fileName)

		sets.SaveSet(Path, set)
	},
}

func init() {
	rootCmd.AddCommand(addRequestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addRequestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addRequestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	addRequestCmd.Flags().StringVarP(&Path, "Path", "P", "", "The path to the set")
        addRequestCmd.MarkFlagRequired("Path")

	addRequestCmd.Flags().StringVarP(&Name, "Name", "N", "", "The name of the request")
	addRequestCmd.MarkFlagRequired("Name")
}
