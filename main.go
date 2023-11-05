package main

import (
	"ffapi/config"
	configmodels "ffapi/config/config-models"
	"ffapi/routes"
	"fmt"
)

func main() {
	serverConfig := config.Cfg.Configs["server"].(*configmodels.ServerConfig)
	addr := fmt.Sprintf(":%d", serverConfig.Port)
	if err := routes.Router.Run(addr); err != nil {
		panic(err)
	}
}
