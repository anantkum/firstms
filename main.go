package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/anantkum/pkg/routes"
	"github.com/anantkum/pkg/utils"
	"github.com/anantkum/pkg/configs"
	"github.com/anantkum/middleware"

)

func main() {
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	utils.StartServer(app)
	/* Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}*/
}
