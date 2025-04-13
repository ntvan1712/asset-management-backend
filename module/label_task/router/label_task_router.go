package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/label_task/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	labelTaskRoute := app.Group("/api/label-tasks")
	labelTaskRoute.Use(middleware.GetAuthMiddleware().AssetManagementAuthorityMiddleware)

	labelTaskController := controller.NewLabelTaskController()
	labelTaskRoute.Post("/presigned-urls", labelTaskController.GetLabelTasksPresignedUrlsHandler)
	labelTaskRoute.Post("/", labelTaskController.CreateLabelTaskHandler)

}
