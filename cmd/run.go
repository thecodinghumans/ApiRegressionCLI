

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/albertapi/AlbertApiCLI/requests"
	"github.com/albertapi/AlbertApiCLI/sets"
	"github.com/albertapi/AlbertApiCLI/findreplaces"
)

var Parallel bool
var PromptEachCall bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the set",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		set := sets.LoadSet(Path)

		var requestsMap = make(map[string]requests.Request)
		for _, requestFileName := range set.Requests {
			requestsMap[requestFileName] = requests.LoadRequest(Path, requestFileName)
		}

		var findReplaceMap = make(map[string]findreplaces.FindReplace)
		for _, findReplaceFileName := range set.FindReplaces {
			findReplaceMap[findReplaceFileName] = findreplaces.LoadFindReplace(Path, findReplaceFileName)
		}

		runSet(set, requestsMap, findReplaceMap)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runCmd.Flags().BoolVar(&Parallel, "Parallel", false, "Run each data row in parallel")

	runCmd.Flags().StringVarP(&Path, "Path", "p", "", "The path to the set")
	runCmd.MarkFlagRequired("Path")

	runCmd.Flags().BoolVar(&PromptEachCall, "PromptEachCall", false, "Prompt the user whether to continue with each run")
}

func makeApiCall(request requests.Request){
	fmt.Println(request)
}

func runSetWithData(set sets.Set, requestsMap map[string]requests.Request, findReplaceMap map[string]findreplaces.FindReplace, dataItem map[string]string, wg *sync.WaitGroup){
	wg.Done()

	for _, item := range set.Requests {
		request := requestsMap[item]
		makeApiCall(request)
	}
}

func runSet(set sets.Set, requestsMap map[string]requests.Request, findReplaceMap map[string]findreplaces.FindReplace){
	var wg sync.WaitGroup

	for _, dataItem := range set.Data {
		if Parallel {
			wg.Add(1)
			go runSetWithData(set, requestsMap, findReplaceMap, dataItem, &wg)
		}else{
			runSetWithData(set, requestsMap, findReplaceMap, dataItem, nil)
		}
	}

	wg.Wait()
}
