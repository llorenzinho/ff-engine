package configmodels

type LogConfig struct {
	Name  string
	Level string `mapstructure:"level"`
}

func (lc *LogConfig) Default() interface{} {
	return &LogConfig{
		Name:  "log",
		Level: "info",
	}
}

func (lc *LogConfig) GetName() string {
	return lc.Name
}
