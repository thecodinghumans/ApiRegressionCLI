package responses

import (
	"time"

	"github.com/thecodinghumans/ApiRegressionCLI/requests"
)

type Response struct {
	OriginalRequest		requests.Request	`json:OriginalRequest`
	ComputedRequest		requests.Request	`json:ComputedRequest`
	Error			string			`json:Error`
	StatusCode		int			`json:StatusCode`
	Headers			map[string][]string	`json:Headers`
	Body			string			`json:Body`
	Timing			int64			`json:Timing`
	MeetsExpectedStatusCode	bool			`json:MeetsExpectedStatusCode`
	MeetsExpectedTiming	bool			`json:MeetsExpectedTiming`
	MeetsExpectedBodyFormat	bool			`json:MeetsExpectedBodyFormat`
	CreateDate		time.Time		`json:CreateDate`
}
