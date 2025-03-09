package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/logger"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/auth/domain/entity"
	"asset_management_backend/module/auth/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
}

func (ac *AuthController) LoginHandler(c *fiber.Ctx) error {

	var loginRequest entity.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		logger.Error("AccountController", "LoginHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(loginRequest); err != nil {
		logger.Error("AccountController", "LoginHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	response, err := ac.authUsecase.Login(c.Context(), loginRequest)
	if err != nil {
		if err == error_app.ErrBadRequest {
			return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Please send username and password"))
		}
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Not found user"))
		}
		if err == error_app.ErrUnauthorized {
			return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Cant login, please check username, password"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func NewAuthController() *AuthController {
	return &AuthController{
		authUsecase: usecase.NewAuthUsecase(),
	}
}
