package envs

import (
	"github.com/thecodinghumans/ApiRegressionCLI/ioUtils"
)

type Env struct {
	Config	map[string]string	`json:Config`
}

func getFileName(path string) string {
        return path + "/env.json"
}

func LoadEnv(path string) Env{
        return ioUtils.Load[Env](getFileName(path))
}

func EnvExists(path string) bool {
        return ioUtils.FileExists(getFileName(path))
}

func SaveEnv(path string, env Env){
        ioUtils.Save[Env](getFileName(path), env)
}

