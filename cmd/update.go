/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thecodinghumans/ApiRegressionCLI/envs"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A command to update existing sets with new fields and files",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !envs.EnvExists(Path) {
			env := envs.Env{
				Config: make(map[string]string),
			}
			envs.SaveEnv(Path, env)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	updateCmd.Flags().StringVarP(&Path, "Path", "P", "", "The path to create the new set (Required)")
        updateCmd.MarkFlagRequired("Path")
}
