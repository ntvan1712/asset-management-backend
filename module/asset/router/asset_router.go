package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/asset/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	assetRoute := app.Group("/api/assets")
	assetController := controller.NewAssetController()
	// Employee
	employeeAuth := middleware.GetAuthMiddleware().EmployeeAuthorityMiddleware
	assetRoute.Get("/borrowed/", employeeAuth, assetController.GetMyBorrowedAssetsHandler)
	
	// Manager
	assetRoute.Use(middleware.GetAuthMiddleware().AssetManagementAuthorityMiddleware)
	assetRoute.Get("/", assetController.SearchByFilterHandler)
	assetRoute.Get("/:asset_id", assetController.GetAssetByIDHandler)
	assetRoute.Put("/:asset_id", assetController.UpdateAssetHandler)
	assetRoute.Post("/", assetController.CreateAssetHandler)
	assetRoute.Post("/presigned-urls", assetController.GetAssetFilesPresignedUrlsHandler)
}
