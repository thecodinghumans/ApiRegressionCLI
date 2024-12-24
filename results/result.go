package results

import (
	"github.com/thecodinghumans/ApiRegressionCLI/responses"
)

type Result struct {
	DataItemKey	string			`json:DataItemKey`
	DataItem	map[string]string	`json:DataItem`
	Responses	[]responses.Response	`json:reponses`
}

