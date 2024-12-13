package cmd

import (
	"fmt"
	"strings"
	"sync"
//	"encoding/gob"
	"bytes"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/tidwall/gjson"
	"github.com/spf13/cobra"
	"github.com/albertapi/AlbertApiCLI/requests"
	"github.com/albertapi/AlbertApiCLI/sets"
	"github.com/albertapi/AlbertApiCLI/findreplaces"
	"github.com/albertapi/AlbertApiCLI/responses"
	"github.com/albertapi/AlbertApiCLI/runresults"
	"github.com/albertapi/AlbertApiCLI/results"
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

		runSet(Name, set, requestsMap, findReplaceMap)
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

	runCmd.Flags().StringVarP(&Name, "Name", "n", "", "A name for this particular run")
}

func makeApiCall(
	request requests.Request,
) (responses.Response, error) {
	var resp responses.Response

	client := &http.Client{Timeout: 10 * time.Minute}

	req, err := http.NewRequest(strings.ToUpper(request.Method), request.Url, bytes.NewBuffer([]byte(request.Body)))
	if err != nil {
		return resp, err
	}

	req.Header.Set("Content-Type", "application/json")

	before := time.Now()

	httpResp, err := client.Do(req)

	timing := time.Since(before)

	if err != nil {
		return resp, err
	}

	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, err
	}

	resp = responses.Response{
		StatusCode: httpResp.StatusCode,
		Headers: httpResp.Header,
		Body: string(body),
		Timing: timing.Milliseconds(),
		MeetsExpectedStatusCode: request.ExpectedStatus == httpResp.StatusCode,
		MeetsExpectedTiming: request.ExpectedTiming >= timing.Milliseconds(),
	}

	return resp, nil
}

func runSetWithData(
	set sets.Set, 
	requestsMap map[string]requests.Request, 
	findReplaceMap map[string]findreplaces.FindReplace, 
	dataItem map[string]string, 
	wg *sync.WaitGroup,
	mx *sync.Mutex,
	runResult *runresults.RunResult,
	dataItemKey string,
) {
	if wg != nil {
		defer wg.Done()
	}

	fmt.Println("Running with data item: " + dataItemKey)

	resps := make([]responses.Response, 0)

	for _, item := range set.Requests {
		request := requestsMap[item]

		newRequest, err := DeepClone[requests.Request](request)
		if err != nil {
			panic(err)
		}

		//newRequest := requestClone.(requests.Request)
		newRequest.Method = GetVal(dataItem, set.Config, request.Method, findReplaceMap, resps)
		newRequest.Url = GetVal(dataItem, set.Config, request.Url, findReplaceMap, resps)
		for key, val := range request.Headers {
			newRequest.Headers[key] = GetVal(dataItem, set.Config, val, findReplaceMap, resps)
		}
		newRequest.Body = GetVal(dataItem, set.Config, request.Body, findReplaceMap, resps)

		resp, err := makeApiCall(newRequest)
		if err != nil {
			resp.Error = err
		}

		resp.OriginalRequest = request
		resp.ComputedRequest = newRequest

		resps = append(resps, resp)

		if !resp.MeetsExpectedStatusCode {
			break
		}
	}

	if mx != nil {
		mx.Lock()
	}

	res := results.Result{
		DataItem: dataItem,
		Responses: resps,
	}

	runResult.Results = append(runResult.Results, res)

	if mx != nil {
		mx.Unlock()
	}

	fmt.Println("Finished with data item: " + dataItemKey)
}

func runSet(
	Name string,
	set sets.Set, 
	requestsMap map[string]requests.Request, 
	findReplaceMap map[string]findreplaces.FindReplace,
) {
	fmt.Println("Running Set: " + set.Name)

	var wg sync.WaitGroup
	var mx sync.Mutex

	runResult := runresults.RunResult{
		Name: Name,
		CreateDate: time.Now().UTC(),
	}

	if !Parallel {
		wg.Add(1)
	}

	for key, dataItem := range set.Data {
		if Parallel {
			wg.Add(1)
			go runSetWithData(set, requestsMap, findReplaceMap, dataItem, &wg, &mx, &runResult, key)
		}else{
			runSetWithData(set, requestsMap, findReplaceMap, dataItem, nil, nil, &runResult, key)
		}
	}

	if !Parallel {
		wg.Done()
	}

	wg.Wait()

	runResult.EndDate = time.Now().UTC()

	infos := runresults.LoadInfo(Path)

	info := runresults.RunResultInfo{
		FileName: runresults.GetFileName(runResult.CreateDate),
		Name: Name,
		CreateDate: runResult.CreateDate,
	}

	infos.Rows = append(infos.Rows, info)

	runresults.SaveRunResult(Path, runResult)

	runresults.SaveInfo(Path, infos)

	fmt.Println("Finished running set: " + set.Name)
}

func GetVal(
	dataItem map[string]string,
	config map[string]string,
	val string,
	findReplaceMap map[string]findreplaces.FindReplace,
	resps []responses.Response,
) string {
	//Build the single  map to use
	mergedMap := MergeMaps(config, dataItem)

	//Ovewrite the map with any findreplace strings
	for _, fr := range findReplaceMap {
		val := fr.Replace

		requestFileName := fr.ReplaceWithRequestFileName
		replaceFrom := fr.ReplaceFrom

		var response responses.Response

		for _, resp := range resps {
			if resp.OriginalRequest.FileName == requestFileName {
				response = resp
			}
		}

		if response.OriginalRequest.FileName != "" {
			switch(replaceFrom) {
				case "Response-Headers":
					val = response.Headers[val][0]
					break
				case "Response-Body":
					val = gjson.Get(response.Body, val).String()
					break
			}
		}

		mergedMap[fr.Find] = val
	}

	return ReplacePlaceholders(val, mergedMap)
}

// ReplacePlaceholders replaces all placeholders in the format {{key}} within a string
// with corresponding values from the replacements map.
func ReplacePlaceholders(input string, replacements map[string]string) string {
	var result strings.Builder
	i := 0

	for i < len(input) {
		// Look for the start of a placeholder
		if i+1 < len(input) && input[i] == '{' && input[i+1] == '{' {
			// Find the end of the placeholder
			end := strings.Index(input[i+2:], "}}")
			if end != -1 {
				// Extract the key between {{ and }}
				key := input[i+2 : i+2+end]
				// Look up the replacement value in the map
				if value, ok := replacements[key]; ok {
					result.WriteString(value)
				} else {
					// If the key is not in the map, write the original placeholder
					result.WriteString("{{" + key + "}}")
				}
				// Move past this placeholder
				i += end + 4
				continue
			}
		}
		// If not a placeholder, just add the character to the result
		result.WriteString(string(input[i]))
		i++
	}

	return result.String()
}

// MergeMaps takes two map[string]string and returns a new merged map.
func MergeMaps(map1, map2 map[string]string) map[string]string {
    // Create a new map to hold the merged values
    mergedMap := make(map[string]string)

    // Copy all elements from map1 to mergedMap
    for key, value := range map1 {
        mergedMap[key] = value
    }

    // Copy all elements from map2 to mergedMap (overwriting any duplicate keys)
    for key, value := range map2 {
        mergedMap[key] = value
    }

    return mergedMap
}

func DeepClone[T any](src T) (T, error) {
	/*
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(src)
	if err != nil {
		return "", err
	}

	gobData := buf.Bytes()

	network2 := bytes.NewBuffer(gobData)
	decoder := gob.NewDecoder(network2)

	var dst string
	err = decoder.Decode(&dst)
	if err != nil {
		return "", err
	}

	return ds, nil
	*/

	var dst T

	jsonString, err := json.Marshal(src)
	if err != nil {
		return dst, err
	}
	err = json.Unmarshal([]byte(jsonString), &dst)
	return dst, err
}
