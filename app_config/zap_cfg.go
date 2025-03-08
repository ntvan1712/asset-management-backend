package app_config

type LoggerConfig struct {
	Development       bool   `mapstructure:"Development"`
	DisableCaller     bool   `mapstructure:"DisableCaller"`
	DisableStacktrace bool   `mapstructure:"DisableStacktrace"`
	Encoding          string `mapstructure:"Encoding"`
	Level             string `mapstructure:"Level"`
	SugarType         bool   `mapstructure:"SugarType"`
}

