package app_config

type RabbitMQConfig struct {
	URI                     string `mapstructure:"URI"`
	LabelTaskQueue          string `mapstructure:"LabelTaskQueue"`
	CompletedLabelTaskQueue string `mapstructure:"VideoTaskQueue"`
}
