package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/asset/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	assetRoute := app.Group("/api/assets")
	assetRoute.Use(middleware.GetAuthMiddleware().AssetManagementAuthorityMiddleware)

	assetController := controller.NewAssetController()
	assetRoute.Get("/:asset_id", assetController.GetAssetByIDHandler)
	assetRoute.Post("/", assetController.CreateAssetHandler)
	assetRoute.Post("/presigned-urls", assetController.GetPresignedUrlsHandler)
}
