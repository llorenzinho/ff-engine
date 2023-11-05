package routes

import (
	"ffapi/config"
	configmodels "ffapi/config/config-models"
	"ffapi/routes/fflag"
	"log"
	"slices"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRoutes() {
	fflag.InitFeatureFlagRoutes(Router)
}

func init() {
	mode := config.Cfg.Configs["server"].(*configmodels.ServerConfig).Mode
	allowed := []string{gin.DebugMode, gin.ReleaseMode, gin.TestMode}
	if !slices.Contains(allowed, mode) {
		log.Printf("Invalid mode %s, defaulting to debug mode", mode)
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	Router = gin.Default()
	InitRoutes()
}
