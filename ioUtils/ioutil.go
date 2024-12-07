package ioUtils

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

func FileExists(fileName string) bool {
        _, err := os.Stat(fileName)
        if err == nil {
                return true
        }
        return false
}

func Load[T any](fileName string) T {
        var t T

        if !FileExists(fileName) {
                return t
        }

        file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
        if err != nil {
                fmt.Println("Error opening file", err)
                return t
        }
        defer file.Close()

        // Read the file's content
        bytes, err := ioutil.ReadAll(file)
        if err != nil {
                fmt.Println("Error reading file", err)
                return t
        }

        // Decode JSON data into the struct
        if err := json.Unmarshal(bytes, &t); err != nil {
                fmt.Println("Error decoding JSON", err)
                return t
        }

        return t
}

func Save[T any](fileName string, t T) {
        file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
        if err != nil {
                fmt.Println("Error opening file", err)
        }
        defer file.Close()

        // Encode the struct as JSON and write >
        encoder := json.NewEncoder(file)
        encoder.SetIndent("", "  ")
        if err := encoder.Encode(t); err != nil{
                fmt.Println("Error encoding JSON", err)
                return
        }
}
