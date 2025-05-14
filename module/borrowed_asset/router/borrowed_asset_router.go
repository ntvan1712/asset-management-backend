package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/borrowed_asset/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// For employee
	borrowedAssetController := controller.NewBorrowedAssetController()

	// For manager and admin
	managerBorrowRoute := app.Group("/api/borrowed-assets/managed/")
	managerBorrowRoute.Use(middleware.GetAuthMiddleware().BorrowManagementAuthorityMiddleware)
	managerBorrowRoute.Delete("/:id", borrowedAssetController.ReturnAssetHandler)
}
