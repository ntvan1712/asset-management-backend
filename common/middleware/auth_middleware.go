package middleware

import (
	"strings"
	"sync"
	"asset_management_backend/app_config"
	"asset_management_backend/common/enums"
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/logger"
	"asset_management_backend/module/auth/domain/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	UserIdFieldName = "user-id"
)

type AuthMiddlewareProvider struct {
	AdminAuthorityMiddleware              fiber.Handler
	AssetManagementAuthorityMiddleware    fiber.Handler
	BorrowManagementAuthorityMiddleware   fiber.Handler
	StatisticalAuthorityMiddleware        fiber.Handler
	CategoryManagementAuthorityMiddleware fiber.Handler
	EmployeeAuthorityMiddleware           fiber.Handler
}

var authMiddlewareProvider *AuthMiddlewareProvider
var initOnce sync.Once

func GetAuthMiddleware() *AuthMiddlewareProvider {
	initOnce.Do(initAuthMiddlewareProvider)
	return authMiddlewareProvider
}

func initAuthMiddlewareProvider() {
	authMiddlewareProvider = &AuthMiddlewareProvider{
		AdminAuthorityMiddleware:              authMiddleware(enums.UserAuthority.Admin),
		AssetManagementAuthorityMiddleware:    authMiddleware(enums.UserAuthority.AssetManagement),
		BorrowManagementAuthorityMiddleware:   authMiddleware(enums.UserAuthority.BorrowManagement),
		StatisticalAuthorityMiddleware:        authMiddleware(enums.UserAuthority.Statistical),
		CategoryManagementAuthorityMiddleware: authMiddleware(enums.UserAuthority.CategoryManagement),
		EmployeeAuthorityMiddleware:           authMiddleware(enums.UserAuthority.Employee),
	}
	logger.Info("[AuthMiddlewareProvider] init AuthMiddlewareProvider")
}

func authMiddleware(userAuthority string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		token, err := verifyAccessToken(c)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(error_app.InvalidAccessTokenErrorResponse("Access token không hợp lệ"))
		}

		if userAuthority != enums.UserAuthority.Employee {
			canAccess, err := usecase.NewAuthUsecase().HasUserAuthority(ctx, token.UserId, userAuthority)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse("Lỗi hệ thống: " + err.Error()))
			}
			if !canAccess {
				return c.Status(fiber.StatusForbidden).JSON(error_app.PermissionDeniedErrorResponse("Không có quyền truy cập"))
			}
		}

		ctx.SetUserValue(UserIdFieldName, token.UserId)
		return c.Next()
	}
}

type JWTAccessTokenEntity struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func verifyAccessToken(c *fiber.Ctx) (*JWTAccessTokenEntity, error) {
	authHeader := c.Get("Authorization")

	if len(strings.TrimSpace(authHeader)) == 0 {
		return nil, error_app.ErrInvalidAccessToken
	}
	accessToken := authHeader[len("Bearer "):]

	if accessToken == "" {
		return nil, error_app.ErrInvalidAccessToken
	}
	token, err := jwt.ParseWithClaims(accessToken, &JWTAccessTokenEntity{}, func(token *jwt.Token) (interface{}, error) {
		return app_config.GetAppConfig().AuthConfig.GetJwtAccessTokenKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTAccessTokenEntity); ok && token.Valid {
		return claims, nil
	} else {
		return nil, error_app.ErrInvalidAccessToken
	}
}
