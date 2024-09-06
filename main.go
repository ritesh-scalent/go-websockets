package main

import (
	"github.com/ritesh-scalent/go-websockets/config"
	httpEngine "github.com/ritesh-scalent/go-websockets/router/http"
)

func main() {
	enviroment := "dev"
	configPath := "path"
	config.Init(&enviroment, &configPath)
	httpEngine.Run("8080")

}
