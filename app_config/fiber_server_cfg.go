package app_config

type FiberServerConfig struct {
	HttpPort      string `mapstructure:"HttpPort"`
	Mode          string `mapstructure:"Mode"`
	BodyLimitInKb int    `mapstructure:"BodyLimitInKb"`
}

func (s *FiberServerConfig) IsProductionMode() bool {
	return s.Mode == "Production"
}
