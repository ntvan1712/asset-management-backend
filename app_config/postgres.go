package app_config

type PostgresConfig struct {
	ConnectionString     string `mapstructure:"ConnectionString"`
	TaskCompletedChannel string `mapstructure:"TaskCompletedChannel"`
}
