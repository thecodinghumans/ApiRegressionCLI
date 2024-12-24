ApiRegressionCLI by The Coding Humans is a command line tool to test APIs for issues by running through one or more calls for one or more scenarios.

Commands
- init: starts a new set at the given path
- addRequest: adds an api request to the set
- addFindReplace: adds logic for doing a find/replace by pulling data from one request and injecting it into another
- run: run the set

Command: init
- Path: the path where you want to store your set
- set.name: the name of the set

Command: addRequest
- Path: the path of the set
- Name: the name of the request
- request.method: the http method for the request (GET, POST, etc)
- request.url: the url for the request
- request.expectedStatus: the expected status code for the responsee (int) ex 200, 401, etc
- request.expectedTiming: the expected timeing for the response (int), ex 1000 where the expecation is that the call will take 1000ms or less

Commmand: addFindReplace
- Path: the path of the set
- Name: a name for this find/replace

Command: run
- Path: the path of the set to run
- Name: a name for this particular run
- Parallel: whether to run each data item in parallel (default is false)
- PromptEachCall: whether to walk through and confirm each call before it is made. This is useful when first running the set to confirm if the requests are valid
- RunEverySeconds: define an int higher than 0 and it will run the set every X seconds

Command: resultSummary
- Path: the path of the set to output the result summary
- Since: the date in the format of YYYY-MM-DD HH:MM:SS to include the result since
