package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/middleware"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/asset_history/domain/entity"
	"asset_management_backend/module/asset_history/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type AssetHistoryController struct {
	historyUsecase usecase.AssetHistoryUsecase
}

func (a *AssetHistoryController) CreateHandler(c *fiber.Ctx) error {
	assetID, err := c.ParamsInt("asset_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Asset Id phải là số nguyên"))
	}

	var request *entity.CreateAssetHistoryRequest
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
	request.CreatorID = &userID
	request.AssetID = &assetID
	response, err := a.historyUsecase.Create(c.Context(), *request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (a *AssetHistoryController) UpdateHandler(c *fiber.Ctx) error {
	historyID, err := c.ParamsInt("asset_history_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Id phải là số nguyên"))
	}

	type updateRequest struct {
		Title   *string `json:"title,omitempty"`
		Content *string `json:"content,omitempty"`
	}
	var request *updateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if request.Title == nil && request.Content == nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Yêu cầu rỗng"))
	}

	if err := a.historyUsecase.UpdateByID(c.Context(), historyID, app_utils.StructToUpdateMap(request)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Update success")
}

func (a *AssetHistoryController) DeleteHandler(c *fiber.Ctx) error {
	historyID, err := c.ParamsInt("asset_history_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Id phải là số nguyên"))
	}

	if err := a.historyUsecase.DeleteByID(c.Context(), historyID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Delete success")
}

func (a *AssetHistoryController) GetByAssetIDHandler(c *fiber.Ctx) error {
	assetID, err := c.ParamsInt("asset_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Asset Id phải là số nguyên"))
	}

	page := c.QueryInt("page", -1)
	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Page không hợp lệ"))
	}

	limit := c.QueryInt("limit", -1)
	if limit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Limit không hợp lệ"))
	}

	result, err := a.historyUsecase.GetByAssetID(c.Context(), assetID, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(result)
}

func NewAssetHistoryController() AssetHistoryController {
	return AssetHistoryController{
		historyUsecase: usecase.NewAssetHistoryUsecase(),
	}
}
