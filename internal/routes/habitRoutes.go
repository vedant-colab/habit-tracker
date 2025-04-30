package routes

import (
	"habit-tracker/internal/dto"
	"habit-tracker/internal/handler"
	"habit-tracker/internal/middleware"

	"github.com/labstack/echo/v4"
)

func HabitRoutes(r *echo.Group) {
	r.POST("/habit", handler.CreateHabitHandler, middleware.JWTMiddleware, middleware.Validate(func() interface{} { return &dto.CreateHabitRequest{} }))
}
