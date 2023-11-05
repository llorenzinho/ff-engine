package config

import (
	configmodels "ffapi/config/config-models"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var (
	Cfg *AppConfig
)

func init() {
	Cfg = GetConfig()
	Cfg.Build()
}

func GetConfig() *AppConfig {
	appcfg := newAppConfig("config", "./config", "APP", "json")
	appcfg.WithConfig(&configmodels.ServerConfig{})
	appcfg.WithConfig(&configmodels.LogConfig{})
	appcfg.WithConfig(&configmodels.K8sClientConfig{})

	return appcfg
}

type IViperConfig interface {
	Default() interface{}
	GetName() string
}

type AppConfig struct {
	Configs map[string]interface{}
}

func newAppConfig(
	configName string, configPath string, envPrefix string, configType string) *AppConfig {
	initConfig(configName, configPath, envPrefix, configType)
	return &AppConfig{
		Configs: make(map[string]interface{}),
	}
}

func (ac *AppConfig) WithConfig(cfg IViperConfig) {
	// Check if config already exists
	if _, ok := ac.Configs[cfg.GetName()]; ok {
		panic(fmt.Errorf("config %s already exists", cfg.GetName()))
	}
	cast, ok := cfg.Default().(IViperConfig)
	if !ok {
		panic(fmt.Errorf("unable to cast config %s to IViperConfig", cfg.GetName()))
	}
	ac.Configs[cast.GetName()] = cast
}

func (ac *AppConfig) Build() {
	for name, cfg := range ac.Configs {
		err := viper.UnmarshalKey(name, cfg)
		if err != nil {
			panic(fmt.Errorf("unable to unmarshal config %s: %w", name, err))
		}
	}
}

func initConfig(configName string, configPath string, envPrefix string, configType string) {
	viper.SetConfigName(configName)
	viper.SetEnvPrefix(envPrefix)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}
	log.Println("Config loaded")
}
