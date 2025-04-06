package app_config

import (
	"fmt"
	"strings"
)

type MinioConfig struct {
	Endpoint        string `mapstructure:"Endpoint"`
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	SecretAccessKey string `mapstructure:"SecretAccessKey"`
	Bucket          string `mapstructure:"Bucket"`
}

func (m *MinioConfig) GetFullAssetUrl(path string) string {
	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("http://%s/%s%s", m.Endpoint, m.Bucket, path)
	}
	return fmt.Sprintf("http://%s/%s/%s", m.Endpoint, m.Bucket, path)
}
