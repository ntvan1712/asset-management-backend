package app_config

type RabbitMQConfig struct {
	URI                     string `mapstructure:"URI"`
	LabelTaskQueue          string `mapstructure:"ImageTaskQueue"`
	CompletedLabelTaskQueue string `mapstructure:"VideoTaskQueue"`
}
