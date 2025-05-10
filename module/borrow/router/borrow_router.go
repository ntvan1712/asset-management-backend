package router

import (
	"asset_management_backend/common/middleware"
	"asset_management_backend/module/borrow/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// For employee
	employeeBorrowRoute := app.Group("/api/borrow-requests")
	employeeBorrowRoute.Use(middleware.GetAuthMiddleware().EmployeeAuthorityMiddleware)

	borrowController := controller.NewBorrowController()
	employeeBorrowRoute.Post("/", borrowController.CreateBorrowRequest)
	employeeBorrowRoute.Delete("/:request_id", borrowController.CancelBorrowRequest)
	employeeBorrowRoute.Get("/", borrowController.GetBorrowRequestsByRequestor)

	// For manager and admin
	managerBorrowRoute := app.Group("/api/borrow-requests/managed/")
	managerBorrowRoute.Use(middleware.GetAuthMiddleware().BorrowManagementAuthorityMiddleware)
	managerBorrowRoute.Get("/", borrowController.GetBorrowRequests)
	managerBorrowRoute.Post("/:request_id/reject", borrowController.RejectBorrowRequestHandler)
	managerBorrowRoute.Post("/:request_id/approve", borrowController.ApproveBorrowRequestHandler)
}
