package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/statistic/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	statisticController := controller.NewStatisticController()
	// For manager and admin
	statisticRoute := app.Group("/api/statistics/")
	statisticRoute.Use(middleware.GetAuthMiddleware().AssetManagementAuthorityMiddleware)
	statisticRoute.Get("/assets/", statisticController.GetAssetStatistic)
}
