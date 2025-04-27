package handler

import (
	"habit-tracker/internal/dto"
	"habit-tracker/internal/service"
	"habit-tracker/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get user by ID
// @Description Get details of a user by their ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id     path     int  true  "User ID"
// @Success 200 {object} UserResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]

func CreateUserHandler(c echo.Context) error {
	userRequest, err := utils.GetDTO[*dto.UserRequest](c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid request context"})
	}
	user, err := service.CreateUserService(*userRequest)
	if err != nil {
		if err.Error() == "user already exists" {
			return c.JSON(http.StatusConflict, echo.Map{
				"success": false,
				"error":   "User already exists",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Something went wrong"})
	}

	c.JSON(http.StatusCreated, echo.Map{
		"success": true,
		"message": "User created successfully",
		"user": echo.Map{
			"username": user.Username,
			"email":    user.Email,
		},
	})
	return nil
}

func LoginHandler(c echo.Context) error {

	loginRequest, err := utils.GetDTO[*dto.LoginRequest](c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid request context"})
	}
	// Check if the user exists and get user from DB
	user, err := service.GetUserByUsername(*loginRequest)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Compare the hashed password
	err = utils.ComparePassword(user.Password, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	// Respond with the generated token
	return c.JSON(http.StatusOK, dto.LoginResponse{Token: token, User: echo.Map{
		"ID":       user.ID,
		"Username": user.Username,
	}})
}
