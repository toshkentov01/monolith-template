package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/sardortoshkentov/mymonolith/api/v1/controllers"
)

// UserRoutes func for describe group of public routes.
func UserRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	// Routes For Post Method:
	route.Post("/register/signup/", controllers.SignUp)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
