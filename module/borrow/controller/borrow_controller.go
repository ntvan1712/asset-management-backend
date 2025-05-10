package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/middleware"
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/borrow/domain/entity"
	"asset_management_backend/module/borrow/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type BorrowController struct {
	borrowUsecase usecase.BorrowUsecase
}

func (b *BorrowController) CreateBorrowRequest(c *fiber.Ctx) error {
	var request *entity.BorrowRequestBody
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	userID, ok := c.Context().UserValue(middleware.UserIdFieldName).(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid user ID"))
	}
	request.RequestorID = &userID
	response, err := b.borrowUsecase.CreateBorrowRequest(c.Context(), *request)
	if err != nil {
		if err == error_app.ErrDuplicateKey {
			return c.Status(fiber.StatusConflict).JSON(error_app.ConflictErrorResponse("Serial number đã tồn tại"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (b *BorrowController) CancelBorrowRequest(c *fiber.Ctx) error {
	requestID, err := c.ParamsInt("request_id")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid request ID: " + err.Error()))
	}
	userID, ok := c.Context().UserValue(middleware.UserIdFieldName).(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid user ID"))
	}
	err = b.borrowUsecase.CancelBorrowRequestByID(c.Context(), requestID, userID)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Yêu cầu không tồn tại hoặc đã bị hủy"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Delete success")
}

func (b *BorrowController) RejectBorrowRequestHandler(c *fiber.Ctx) error {
	type rejectBody struct {
		ResponseDescription *string `json:"response_description,omitempty"`
	}
	var body rejectBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}

	requestID, err := c.ParamsInt("request_id")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid request ID: " + err.Error()))
	}

	userID, ok := c.Context().UserValue(middleware.UserIdFieldName).(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid user ID"))
	}

	err = b.borrowUsecase.RejectBorrowRequest(c.Context(), requestID, userID, body.ResponseDescription)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Yêu cầu không tồn tại hoặc đã bị hủy"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Reject success")
}

func (b *BorrowController) ApproveBorrowRequestHandler(c *fiber.Ctx) error {
	type approveBody struct {
		AssetID             int     `json:"asset_id" validate:"required"`
		ResponseDescription *string `json:"response_description,omitempty"`
	}
	var body approveBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	requestID, err := c.ParamsInt("request_id")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid request ID: " + err.Error()))
	}

	userID, ok := c.Context().UserValue(middleware.UserIdFieldName).(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid user ID"))
	}

	err = b.borrowUsecase.ApproveBorrowRequest(c.Context(), requestID, userID, body.AssetID, body.ResponseDescription)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Yêu cầu không tồn tại hoặc đã bị hủy"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Approve success")
}

func (b *BorrowController) GetBorrowRequestsByRequestor(c *fiber.Ctx) error {

	type getByRequestorIDQuery struct {
		RequestorID int `query:"requestor_id" validate:"required"`

		sharedmodel.PaginateQuery
	}
	var query getByRequestorIDQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	response, err := b.borrowUsecase.GetBorrowRequestsByRequestorID(c.Context(), query.RequestorID, query.PaginateQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (b *BorrowController) GetBorrowRequests(c *fiber.Ctx) error {
	paginateQuery, err := sharedmodel.GetPaginateQuery(c)
	if err != nil {
		return err
	}
	response, err := b.borrowUsecase.GetPendingRequests(c.Context(), *paginateQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func NewBorrowController() *BorrowController {
	return &BorrowController{
		borrowUsecase: usecase.NewBorrowRequestUsecase(),
	}
}
