package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/sardortoshkentov/mymonolith/api/v1/models"
	exchangemodels "gitlab.com/sardortoshkentov/mymonolith/exchange_models"
	"gitlab.com/sardortoshkentov/mymonolith/service"
)

// SignUp method to create a new user.
// @Description Create a new user.
// @Summary creates a new user
// @Tags register
// @Accept json
// @Produce json
// @Param register body models.SignUp true "register"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/register/signup/ [post]
func SignUp(c *fiber.Ctx) error {
	var (
		body models.SignUp
	)

	err := c.BodyParser(&body)
	if err != nil {
		log.Println("Error parsing body: ", err)
		c.Status(http.StatusBadRequest).JSON(models.Response{
			Error: true,
			Data: models.Error{
				Status:  "Bad Request",
				Message: "Failed to Parse a body",
			},
		})
		return err
	}

	userID, err := uuid.NewRandom()
	if err != nil {
		log.Println("Error while generating user id")
		return c.Status(http.StatusBadRequest).JSON(models.Response{
			Error: true,
			Data: models.Error{
				Status:  "Bad Request",
				Message: "Failed to generate user id",
			},
		})
	}

	_, err = service.User().CreateUser(context.Background(), &exchangemodels.CreateUserModel{
		ID:       userID.String(),
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		log.Println("Error while creating user in api", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(models.Response{
			Error: true,
			Data: models.Error{
				Status:  "Internal Server Error",
				Message: err.Error(),
			},
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Error: false,
		Data: models.SuccessMessage{
			Success: true,
		},
	})
}
