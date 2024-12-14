package requests

import (
	"github.com/thecodinghumans/ApiRegressionCLI/ioUtils"
)

type Request struct{
	FileName		string			`json:FileName`
	Name			string 			`json:Name`
	Method			string			`json:Method`
	Url			string			`json:Url`
	Headers			map[string]string	`json:Headers`
	Body			string			`json:Body`
	ExpectedStatus		int			`json:ExpectedStatus`
	ExpectedTiming		int64			`json:ExpectedTiming`
	ExpectedBodyFormat	string			`json:ExpectedBodyFormat`
}

func getFileName(path string, fileName string) string {
	return path + "/Requests/" + fileName
}

func LoadRequest(path string, fileName string) Request {
	return ioUtils.Load[Request](getFileName(path, fileName))
}

func RequestExists(path string, fileName string) bool {
	return ioUtils.FileExists(getFileName(path, fileName))
}

func SaveRequest(path string, fileName string, request Request){
	ioUtils.Save[Request](getFileName(path, fileName), request)
}
