package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/category/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	categoryRoute := app.Group("/api/categories")

	categoryController := controller.NewCategoryController()
	categoryRoute.Get("/", categoryController.GetAllHandler)

	categoryRoute.Use(middleware.GetAuthMiddleware().CategoryManagementAuthorityMiddleware)

	assetQualityController := controller.NewAssetQualityController()
	assetQualityRoute := categoryRoute.Group("/asset-qualities")
	assetQualityRoute.Get("/", assetQualityController.GetAllHandler)
	assetQualityRoute.Post("/", assetQualityController.CreateHandler)
	assetQualityRoute.Delete("/:asset_quality_id", assetQualityController.DeleteByIDHandler)
	assetQualityRoute.Put("/:asset_quality_id", assetQualityController.UpdateHandler)

	assetTypeController := controller.NewAssetTypeController()
	assetTypeRoute := categoryRoute.Group("/asset-types")
	assetTypeRoute.Get("/", assetTypeController.GetAllHandler)
	assetTypeRoute.Post("/", assetTypeController.CreateHandler)
	assetTypeRoute.Delete("/:asset_type_id", assetTypeController.DeleteByIDHandler)
	assetTypeRoute.Put("/:asset_type_id", assetTypeController.UpdateHandler)

	locationController := controller.NewLocationController()
	locationRoute := categoryRoute.Group("/locations")
	locationRoute.Get("/", locationController.GetAllHandler)
	locationRoute.Post("/", locationController.CreateHandler)
	locationRoute.Delete("/:location_id", locationController.DeleteByIDHandler)
	locationRoute.Put("/:location_id", locationController.UpdateHandler)

	currencyController := controller.NewCurrencyController()
	currencyRoute := categoryRoute.Group("/currencies")
	currencyRoute.Get("/", currencyController.GetAllHandler)
}
