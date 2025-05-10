package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/transfer/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// For employee
	employeeTransferRoute := app.Group("/api/transfer-requests/")
	employeeTransferRoute.Use(middleware.GetAuthMiddleware().EmployeeAuthorityMiddleware)

	transferController := controller.NewTransferController()
	employeeTransferRoute.Get("/", transferController.GetMyTransferRequestsHandler)
	employeeTransferRoute.Post("/:request_id/reject", transferController.RejectTransferRequestHandler)
	employeeTransferRoute.Post("/:request_id/approve", transferController.ApproveTransferRequestHandler)

	// For manager and admin
	managerTransferRoute := app.Group("/api/transfer-requests/managed/")
	managerTransferRoute.Use(middleware.GetAuthMiddleware().BorrowManagementAuthorityMiddleware)
	managerTransferRoute.Get("/", transferController.GetAllHandler)
	managerTransferRoute.Delete("/:request_id/", transferController.CancelTransferRequestHandler)
}
