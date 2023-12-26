package server

import (
	"FleetManagerAPI/config"
	"FleetManagerAPI/platform/database"
	"FleetManagerAPI/platform/logger"

	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func Serve() {
	appCfg := config.AppCfg()

	logger.SetUpLogger() // Set up logger

	logr := logger.GetLogger()

	if err := database.Connect(config.DBCfg()); err != nil {
		logr.Panicf("failed database setup. error: %v", err)
	}

	fiberCfg := config.FiberConfig()

	app := fiber.New(fiberCfg)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		logr.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}
}