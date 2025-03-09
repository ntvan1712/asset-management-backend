package router

import (
	"asset_management_backend/module/auth/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	accountController := controller.NewAuthController()

	authRoute := app.Group("/api/auth")
	authRoute.Post("/login", accountController.LoginHandler)
}
