package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	pkgLogger "gitlab.com/sardortoshkentov/mymonolith/pkg/logger"
	"gitlab.com/sardortoshkentov/mymonolith/config"
	"gitlab.com/sardortoshkentov/mymonolith/pkg/middleware"
	"gitlab.com/sardortoshkentov/mymonolith/pkg/routes"
	"gitlab.com/sardortoshkentov/mymonolith/pkg/utils"
	_ "gitlab.com/sardortoshkentov/mymonolith/api/docs"
)

var (
	fiberConfig = config.FiberConfig()
	appConfig   = config.Config()
)

// @title MyMonolith API
// @version 0.1
// @description This is an auto-generated API Docs for MyMonolith.
// @termsOfService http://swagger.io/terms/
// @contact.name MyMonolith
// @contact.email info@mymonolith.uz
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	log.Println("Started")
	// Define a new Fiber app with config.
	app := fiber.New(fiberConfig)

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	middleware.FiberMiddleware(app)

	_, adapter := utils.InitStorage(config.Config())

	jwtRoleAuthorizer, err := middleware.NewJWTRoleAuthorizer(appConfig, adapter)
	if err != nil {
		log.Fatal("Could not initialize JWT Role Authorizer")
	}

	// Connection to a User Service
	cfg := config.Config()

	log := pkgLogger.New(cfg.LogLevel, "mymonolith")
	defer pkgLogger.Cleanup(log)

	// _ = service.NewUserService(log)

	app.Use(middleware.NewAuthorizer(jwtRoleAuthorizer))
	routes.SwaggerRoute(app)
	routes.UserRoutes(app)

	// Start server (with or without graceful shutdown).
	if config.Config().Environment == "develop" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
