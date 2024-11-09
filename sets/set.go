package sets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Set struct {
	Name      	string 		`json:"name"`
	Config    	string 		`json:"config"`
	Data      	string 		`json:"data"`
	DataLabel 	string 		`json:"dataLabel"`
	Requests	[]string	`json:Requests`
	FindReplaces	[]string	`json:FindReplaces`
}

func LoadSet(path string) Set{
	var set Set

	file, err := os.OpenFile(path + "/set.json", os.O_RDONLY, os.ModePerm)
        if err != nil {
                fmt.Println("Error opening file", err)
                return set
        }
        defer file.Close()

        // Read the file's content
        bytes, err := ioutil.ReadAll(file)
        if err != nil {
                fmt.Println("Error reading file", err)
                return set
        }

        // Decode JSON data into the struct
        if err := json.Unmarshal(bytes, &set); err != nil{
                fmt.Println("Error decoding JSON", err)
                return set
        }

	return set
}

func SaveSet(path string, set Set){
	file, err := os.OpenFile(path + "/set.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	// Encode the struct as JSON and write >
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") 
	if err := encoder.Encode(set); err != nil{
	        fmt.Println("Error encoding JSON", err)
	        return
	}
}
