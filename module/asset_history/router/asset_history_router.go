package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/asset_history/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	historyRoute := app.Group("/api/assets/:asset_id/histories")
	historyRoute.Use(middleware.GetAuthMiddleware().AssetManagementAuthorityMiddleware)

	assetController := controller.NewAssetHistoryController()
	historyRoute.Get("/", assetController.GetByAssetIDHandler)
	historyRoute.Post("/", assetController.CreateHandler)
	historyRoute.Put("/:asset_history_id", assetController.UpdateHandler)
	historyRoute.Delete("/:asset_history_id", assetController.DeleteHandler)
}
