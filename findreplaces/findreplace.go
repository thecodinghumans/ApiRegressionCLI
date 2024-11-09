package findreplaces

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FindReplace struct {
	FileName	string `json:FileName`
	Name		string `json:Name`
	FindType	string `json:FindType`
	FindBy		string `json:FindBy`
	ReplaceType	string `json:ReplaceType`
	ReplaceBy	string `json:ReplaceBy`
}

func LoadFindReplace(path string, fileName string) FindReplace{
	var findReplace FindReplace

	if !FindReplaceExists(path, fileName) {
		return findReplace
	}

	file, err := os.OpenFile(path + "/FindReplaces/" + fileName, os.O_RDONLY, os.ModePerm)
        if err != nil {
                fmt.Println("Error opening file", err)
                return findReplace
        }
        defer file.Close()

        // Read the file's content
        bytes, err := ioutil.ReadAll(file)
        if err != nil {
                fmt.Println("Error reading file", err)
                return findReplace
        }

        // Decode JSON data into the struct
        if err := json.Unmarshal(bytes, &findReplace); err != nil{
                fmt.Println("Error decoding JSON", err)
                return findReplace
        }

        return findReplace
}

func FindReplaceExists(path string, fileName string) bool {
	_, err := os.Stat(path + "/FindReplaces/" + fileName)
	if err == nil {
		return true
	}
	return false
}

func SaveFindReplace(path string, fileName string, findReplace FindReplace){
        file, err := os.OpenFile(path + "/FindReplaces/" + fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
        if err != nil {
                fmt.Println("Error opening file", err)
        }
        defer file.Close()

        // Encode the struct as JSON and write >
        encoder := json.NewEncoder(file)
        encoder.SetIndent("", "  ")
        if err := encoder.Encode(findReplace); err != nil{
                fmt.Println("Error encoding JSON", err)
                return
        }
}
