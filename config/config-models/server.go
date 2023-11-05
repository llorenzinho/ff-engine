package configmodels

type ServerConfig struct {
	Name string
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

func (sc *ServerConfig) Default() interface{} {
	return &ServerConfig{
		Name: "server",
		Mode: "debug",
		Port: 8080,
	}
}

func (sc *ServerConfig) GetName() string {
	return sc.Name
}
