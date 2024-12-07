package runresults

import (
	"time"
	"github.com/albertapi/AlbertApiCLI/results"
	"github.com/albertapi/AlbertApiCLI/ioUtils"
)

type RunResult struct {
	Name		string			`json:Name`
	CreateDate	time.Time		`json:CreateDate`
	EndDate		time.Time		`json:EndDate`
	Results		[]results.Result	`json:Results`
}

type RunResultInfo struct {
	FileName	string		`json:FileName`
	Name		string		`json:Name`
	CreateDate	time.Time	`json:CreateDate`
}

type RunResultInfos struct {
	Rows	[]RunResultInfo	`json:Rows`
}

func GetFileName(date time.Time) string {
	return date.Format("2006-01-02 15:04:05.000") + ".json"
}

func getFullFileName(path string, fileName string) string {
	return path + "/RunResults/" + fileName
}

func LoadRunResult(path string, fileName string) RunResult {
	return ioUtils.Load[RunResult](getFullFileName(path, fileName))
}

func SaveRunResult(path string, runResult RunResult) {
	ioUtils.Save[RunResult](getFullFileName(path, GetFileName(runResult.CreateDate)), runResult)
}

func getInfoFullFileName(path string) string {
	return path + "/RunResults/info.json"
}

func LoadInfo(path string) RunResultInfos {
	return ioUtils.Load[RunResultInfos](getInfoFullFileName(path))
}

func SaveInfo(path string, info RunResultInfos) {
	ioUtils.Save[RunResultInfos](getInfoFullFileName(path), info)
}
