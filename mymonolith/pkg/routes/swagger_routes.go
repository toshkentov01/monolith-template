package routes

import (
	"github.com/gofiber/fiber/v2"
	// "log"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(app *fiber.App) {
	// log.Fatal(app.Listen(":8000"))

	// Create routes group.
	route := app.Group("/swagger/")

	// Routes for GET method:
	route.Get("*", swagger.HandlerDefault) // get one user by ID
}
