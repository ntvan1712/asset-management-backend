package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/middleware"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/label_task/domain/entity"
	"asset_management_backend/module/label_task/domain/usecase"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LabelTaskController struct {
	labelTaskUsecase usecase.LabelTaskUsecase
}

func (a *LabelTaskController) GetLabelTasksPresignedUrlsHandler(c *fiber.Ctx) error {
	var fileNames []string
	if err := c.BodyParser(&fileNames); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	response, err := a.labelTaskUsecase.GetLabelTasksPresignedUrls(c.Context(), fileNames)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (a *LabelTaskController) CreateLabelTaskHandler(c *fiber.Ctx) error {
	var request *entity.LabelTaskRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	rawCtx := c.Context()

	userID, ok := rawCtx.UserValue(middleware.UserIdFieldName).(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Invalid user ID"))
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	rawCtx.SetBodyStreamWriter(func(w *bufio.Writer) {
		fmt.Fprintf(w, "data: %s\n\n", "start")
		w.Flush()

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		responseCh, err := a.labelTaskUsecase.CreateTask(timeoutCtx, *request, userID)
		if err != nil {
			fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
			w.Flush()
			return
		}

		for {
			select {
			case <-timeoutCtx.Done():
				fmt.Fprintf(w, "event: error\ndata: %s\n\n", "timeout or client disconnected")
				w.Flush()
				return
			case task, ok := <-responseCh:
				if !ok {
					return
				}
				if task == nil {
					continue
				}
				if task.Error != nil {
					fmt.Fprintf(w, "event: error\ndata: %s\n\n", task.Error)
				} else {
					labelTaskData, _ := json.Marshal(task.LabelTask)
					fmt.Fprintf(w, "data: %s\n\n", labelTaskData)
				}
				w.Flush()
			}
		}

	})

	return nil
}

func NewLabelTaskController() *LabelTaskController {
	return &LabelTaskController{
		labelTaskUsecase: usecase.NewLabelTaskUsecase(),
	}
}
