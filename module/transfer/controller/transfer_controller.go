package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/middleware"
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/transfer/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type TransferController struct {
	transferUsecase usecase.TransferUsecase
}

func (b *TransferController) GetAllHandler(c *fiber.Ctx) error {
	paginateQuery, err := sharedmodel.GetPaginateQuery(c)
	if err != nil {
		return err
	}
	response, err := b.transferUsecase.GetAll(c.Context(), *paginateQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (b *TransferController) GetMyTransferRequestsHandler(c *fiber.Ctx) error {
	userID, ok := c.Context().UserValue(middleware.UserIdFieldName).(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid user ID"))
	}
	paginateQuery, err := sharedmodel.GetPaginateQuery(c)
	if err != nil {
		return err
	}
	response, err := b.transferUsecase.GetMyTransferRequests(c.Context(), userID, *paginateQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (t *TransferController) ApproveTransferRequestHandler(c *fiber.Ctx) error {
	type approveBody struct {
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

	err = t.transferUsecase.ApproveTransferRequest(c.Context(), requestID, userID, body.ResponseDescription)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Yêu cầu không tồn tại hoặc đã bị hủy"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Approve success")
}

func (t *TransferController) RejectTransferRequestHandler(c *fiber.Ctx) error {
	type rejectBody struct {
		ResponseDescription *string `json:"response_description,omitempty"`
	}
	var body rejectBody

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

	err = t.transferUsecase.RejectTransferRequest(c.Context(), requestID, userID, body.ResponseDescription)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Yêu cầu không tồn tại hoặc đã bị hủy"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Reject success")
}

func (t *TransferController) CancelTransferRequestHandler(c *fiber.Ctx) error {

	requestID, err := c.ParamsInt("request_id")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid request ID: " + err.Error()))
	}

	userID, ok := c.Context().UserValue(middleware.UserIdFieldName).(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid user ID"))
	}

	err = t.transferUsecase.CancelTransferRequest(c.Context(), requestID, userID)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Yêu cầu không tồn tại hoặc đã bị hủy"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Cancel success")
}

func NewTransferController() *TransferController {
	return &TransferController{
		transferUsecase: usecase.NewTransferUsecase(),
	}
}
