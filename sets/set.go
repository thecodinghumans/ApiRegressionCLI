package sets

import (
	"github.com/albertapi/AlbertApiCLI/ioUtils"
)

type Set struct {
	Name      	string 				`json:"name"`
	Config    	map[string]string		`json:"config"`
	Data      	map[string]map[string]string 	`json:"data"`
	Requests	[]string			`json:Requests`
	FindReplaces	[]string			`json:FindReplaces`
}

func getFileName(path string) string {
	return path + "/set.json"
}

func LoadSet(path string) Set{
	return ioUtils.Load[Set](getFileName(path))
}

func SetExists(path string) bool {
	return ioUtils.FileExists(getFileName(path))
}

func SaveSet(path string, set Set){
	ioUtils.Save[Set](getFileName(path), set)
}
