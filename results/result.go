package results

import (
	"github.com/albertapi/AlbertApiCLI/responses"
)

type Result struct {
	DataItem	map[string]string	`json:DataItem`
	Responses	[]responses.Response	`json:reponses`
}

