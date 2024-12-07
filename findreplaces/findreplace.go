package findreplaces

import (
	"github.com/albertapi/AlbertApiCLI/ioUtils"
)

type FindReplace struct {
	FileName			string	`json:FileName`
	Name				string	`json:Name`
	Find				string	`json:Find`
	ReplaceWithRequestFileName	string	`json:ReplaceWithResponseId`
	ReplaceFrom			string	`json:ReplaceFrom`
	Replace				string	`json:Replace`
}

func getFileName(path string, fileName string) string {
	return path + "/FindReplaces/" + fileName
}

func LoadFindReplace(path string, fileName string) FindReplace {
	return ioUtils.Load[FindReplace](getFileName(path, fileName))
}

func FindReplaceExists(path string, fileName string) bool {
	return ioUtils.FileExists(getFileName(path, fileName))
}

func SaveFindReplace(path string, fileName string, findReplace FindReplace){
        ioUtils.Save[FindReplace](getFileName(path, fileName), findReplace)
}
