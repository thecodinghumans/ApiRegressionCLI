package responses

import (
	"time"

	"github.com/albertapi/AlbertApiCLI/requests"
)

type Response struct {
	OriginalRequest		requests.Request			`json:OriginalRequest`
	ComputedRequest		requests.Request			`json:ComputedRequest`
	StatusCode		int			`json:StatusCode`
	Headers			map[string]string	`json:Headers`
	Body			string			`json:Body`
	Timing			uint64			`json:Timing`
	MeetsExpectedStatusCode	bool			`json:MeetsExpectedStatusCode`
	MeetsExpectedTiming	bool			`json:MeetsExpectedTiming`
	MeetsExpectedBodyFormat	bool			`json:MeetsExpectedBodyFormat`
	CreateDate		time.Time		`json:CreateDate`
}
