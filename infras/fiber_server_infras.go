package infras

import (
	// "com.pegatech.faceswap/common/middleware"
	"asset_management_backend/app_config"
	"asset_management_backend/common/logger"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
)

var fiberApp *fiber.App
var initFiberOnce sync.Once

func GetFiberApp() *fiber.App {
	initFiberOnce.Do(initFiberApp)
	return fiberApp
}

func initFiberApp() {
	serverCfg := app_config.GetAppConfig().FiberServerConfig
	fiberApp = fiber.New(
		fiber.Config{
			// ErrorHandler: middleware.BodyLimitHandler,
			BodyLimit: serverCfg.BodyLimitInKb * 1024,
		},
	)

	fiberLogger := fiber_logger.New(
		fiber_logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		},
	)

	fiberApp.Use(fiberLogger)
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Cho phép tất cả các domain
		AllowHeaders: "Content-Type, Authorization",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	logger.Info("[ServerInfras] Init Server fiber app")
}
