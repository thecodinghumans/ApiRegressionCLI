package results

import (
	"github.com/thecodinghumans/ApiRegressionCLI/responses"
)

type Result struct {
	DataItem	map[string]string	`json:DataItem`
	Responses	[]responses.Response	`json:reponses`
}

