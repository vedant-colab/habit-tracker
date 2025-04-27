package routes

import (
	"habit-tracker/internal/dto"
	"habit-tracker/internal/handler"
	"habit-tracker/internal/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(r *echo.Group) {

	r.POST("/users", handler.CreateUserHandler, middleware.Validate(func() interface{} { return &dto.UserRequest{} }))
	r.POST("/login", handler.LoginHandler, middleware.Validate(func() interface{} { return &dto.LoginRequest{} }))
}
