package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Request struct{
	FileName		string		`json:FileName`
	Name			string 		`json:Name`
	Method			string		`json:Method`
	Url			string		`json:Url`
	Headers			[]Header	`json:Headers`
	Body			string		`json:Body`
	ExpectedStatus		int		`json:ExpectedStatus`
	ExpectedTiming		int		`json:ExpectedTiming`
	ExpectedBodyFormat	string		`json:ExpectedBodyFormat`
}

type Header struct{
	Key	string 	`json:Key`
	Value	string	`json:Value`
}

func LoadRequest(path string, fileName string) Request {
	var request Request

	file, err := os.OpenFile(path + "/Requests/" + fileName, os.O_RDONLY, os.ModePerm)
        if err != nil {
                fmt.Println("Error opening file", err)
                return request
        }
        defer file.Close()

        // Read the file's content
        bytes, err := ioutil.ReadAll(file)
        if err != nil {
                fmt.Println("Error reading file", err)
                return request
        }

        // Decode JSON data into the struct
        if err := json.Unmarshal(bytes, &request); err != nil{
                fmt.Println("Error decoding JSON", err)
                return request
        }

	return request
}

func SaveRequest(path string, fileName string, request Request){
	file, err := os.OpenFile(path + "/Requests/" + fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
        if err != nil {
                fmt.Println("Error opening file", err)
        }
        defer file.Close()

        // Encode the struct as JSON and write >
        encoder := json.NewEncoder(file)
        encoder.SetIndent("", "  ")
        if err := encoder.Encode(request); err != nil{
                fmt.Println("Error encoding JSON", err)
                return
        }
}
