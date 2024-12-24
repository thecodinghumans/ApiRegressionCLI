


/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"time"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thecodinghumans/ApiRegressionCLI/runresults"
)

var Since string

// resultSummaryCmd represents the resultSummary command
var resultSummaryCmd = &cobra.Command{
	Use:   "resultSummary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pulling together the results since " + Since)

		layout := "2006-01-02 15:04:05"
		parsedTime, err := time.Parse(layout, Since)
		if err != nil {
			fmt.Println(err)
			return
		}

		var runCount int
		totals := make(map[string]requestResults)
		details := make(map[string]map[string]requestResults)

		info := runresults.LoadInfo(Path)

		for _, row := range info.Rows {
			if row.CreateDate.After(parsedTime) {
				runResult := runresults.LoadRunResult(Path, row.FileName)
				for _, result:= range runResult.Results {
					for _, resp := range result.Responses {
						runCount++
						t, exists := totals[resp.OriginalRequest.Name]
						if !exists {
							t = requestResults{}
						}
						if resp.MeetsExpectedStatusCode {
							t.TotalMeetsExpectedStatusCode++
						}
						if resp.MeetsExpectedTiming {
							t.TotalMeetsExpectedTiming++
						}
						if resp.MeetsExpectedBodyFormat {
							t.TotalMeetsExpectedBodyFormat++
						}
						totals[resp.OriginalRequest.Name] = t


						detail, detailExists := details[result.DataItemKey]
						if !detailExists {
							detail = make(map[string]requestResults)
						}
						innerDetail, innerDetailExists := detail[resp.OriginalRequest.Name]
						if !innerDetailExists {
							innerDetail = requestResults{}
						}
						if resp.MeetsExpectedStatusCode {
							innerDetail.TotalMeetsExpectedStatusCode++
						}
						if resp.MeetsExpectedTiming {
							innerDetail.TotalMeetsExpectedTiming++
						}
						if resp.MeetsExpectedBodyFormat {
							innerDetail.TotalMeetsExpectedBodyFormat++
						}
						detail[resp.OriginalRequest.Name] = innerDetail
						details[result.DataItemKey] = detail
					}
				}
			}
		}

		fmt.Println("Run count: " + strconv.Itoa(runCount))

		fmt.Println("")

		jsonBytes, err := json.MarshalIndent(totals, "", " ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonBytes))

		fmt.Println("")

		detailJsonBytes, err := json.MarshalIndent(details, "", " ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(detailJsonBytes))
	},
}

func init() {
	rootCmd.AddCommand(resultSummaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resultSummaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resultSummaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	resultSummaryCmd.Flags().StringVar(&Path, "path", "", "Path to the set")
	resultSummaryCmd.Flags().StringVar(&Since, "since", "", "Results since when?")
}

type requestResults struct {
	TotalMeetsExpectedStatusCode	int
	TotalMeetsExpectedTiming	int
	TotalMeetsExpectedBodyFormat	int
}
