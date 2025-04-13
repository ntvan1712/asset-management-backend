package main

import (
	"asset_management_backend/app_config"
	"asset_management_backend/cmd"
	"asset_management_backend/common/logger"
	"asset_management_backend/infras"
	assetRouter "asset_management_backend/module/asset/router"
	categoryRouter "asset_management_backend/module/category/router"
	labelTaskRouter "asset_management_backend/module/label_task/router"
	"log"
)

func main() {
	cmd.Execute()
	// Read config file
	app_config.LoadAppConfig(cmd.ConfigFileName)

	if err := logger.InitAppLogger(); err != nil {
		log.Fatal("[Main] Failed to create ZapLogger", err)
	}
	defer logger.Sync()
	logger.Info("[Main] Complete ZapLogger configuration")

	employeeServiceConn := infras.GetEmployeeServiceConn()
	defer employeeServiceConn.Close()

	defer infras.GetRabbitMQProvider().Channel.Close()

	fiberApp := infras.GetFiberApp()

	categoryRouter.Setup(fiberApp)
	assetRouter.Setup(fiberApp)
	labelTaskRouter.Setup(fiberApp)

	// taskRepo := repository.NewLabelTaskRepository()
	// for _ = range 2 {
	// 	taskRepo.CreateTask(
	// 		context.Background(),
	// 		entity.LabelTaskRequest{
	// 			LabelImagePath: "labels/17.jpg",
	// 			TaskType:       "search_asset",
	// 		},
	// 		1001,
	// 	)
	// }
	// go func () {
	// 	datasource.GetTaskStream()
	// }()

	fiberAppErr := fiberApp.Listen(app_config.GetAppConfig().FiberServerConfig.HttpPort)
	if fiberAppErr != nil {
		logger.Fatal("[Main] Failed to start server: " + fiberAppErr.Error())
	}

}
