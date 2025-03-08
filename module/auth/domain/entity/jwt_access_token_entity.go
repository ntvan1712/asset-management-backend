package entity

import (
	// "asset_management_backend/app_config"
	// "asset_management_backend/common/error_app"
	// "time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAccessTokenEntity struct {
	EmployeeID int    `json:"employee_id"`
	Username   string `json:"username"`
	jwt.RegisteredClaims
}

// func GenerateJWTAccessToken(employeeID int, username string) (string, error) {
// 	authConfig := app_config.GetAppConfig().AuthConfig
// 	tokenAliveTime := authConfig.JwtAccessTokenExpireTimeInMinutes
// 	claims := &JWTAccessTokenEntity{
// 		EmployeeID: employeeID,
// 		Username:   username,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenAliveTime) * time.Minute)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(authConfig.GetJwtAccessTokenKey())
// }

// func VerifyJWT(tokenString string) (*JWTAccessTokenEntity, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &JWTAccessTokenEntity{}, func(token *jwt.Token) (interface{}, error) {
// 		return app_config.GetAppConfig().AuthConfig.GetJwtAccessTokenKey(), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(*JWTAccessTokenEntity); ok && token.Valid {
// 		return claims, nil
// 	} else {
// 		return nil, error_app.ErrInvalidAccessToken
// 	}
// }
