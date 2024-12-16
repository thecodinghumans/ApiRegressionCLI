full_build:
	go build
	GOOS=windows GOARCH=amd64 go build -o ApiRegressionCLI.exe main.go
