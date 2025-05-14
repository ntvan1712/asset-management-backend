package controller

import (
	"asset_management_backend/common/enums"
	"asset_management_backend/common/error_app"
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/statistic/data/data_source"
	"asset_management_backend/module/statistic/data/model"

	"github.com/gofiber/fiber/v2"
)

type StatisticController struct {
	source datasource.StatisticDataSource
}

func (s *StatisticController) GetAssetStatistic(c *fiber.Ctx) error {
	category := c.Query("by_category")
	var response *model.CategoryStatisticResponse
	var err error
	if category == enums.AssetStatisticEnum.AssetType {
		response, err = s.source.GetAssetStatisticByType(c.Context())
	} else if category == enums.AssetStatisticEnum.AssetQuality {
		response, err = s.source.GetAssetStatisticByQuality(c.Context())
	} else if category == enums.AssetStatisticEnum.Location {
		response, err = s.source.GetAssetStatisticByLocation(c.Context())
	} else if category == enums.AssetStatisticEnum.Status {
		response, err = s.source.GetAssetStatisticByStatus(c.Context())
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Category không hợp lệ"))
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func NewStatisticController() *StatisticController {
	return &StatisticController{
		source: datasource.NewStatisticDataSource(infras.GetDbProvider()),
	}
}
