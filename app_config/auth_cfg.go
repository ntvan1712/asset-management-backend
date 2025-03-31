package app_config

type AuthConfig struct {
	jwtAccessTokenKeyString           string `mapstructure:"JwtAccessTokenKeyString"`
	JwtAccessTokenExpireTimeInMinutes int    `mapstructure:"JwtAccessTokenExpireTimeInMinutes"`
}

// Token key đảm bảo không thể tạo ngược jwt token từ data object
func (a *AuthConfig) GetJwtAccessTokenKey() []byte {
	return []byte(a.jwtAccessTokenKeyString)
}
