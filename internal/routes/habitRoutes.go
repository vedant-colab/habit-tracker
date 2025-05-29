package routes

import (
	"habit-tracker/internal/dto"
	"habit-tracker/internal/handler"
	"habit-tracker/internal/middleware"

	"github.com/labstack/echo/v4"
)

func HabitRoutes(r *echo.Group) {
	r.POST("/habit", handler.CreateHabitHandler, middleware.JWTMiddleware, middleware.Validate(func() interface{} { return &dto.CreateHabitRequest{} }))
	r.PUT("/habit/:habitid", handler.UpdateHabitHandler, middleware.JWTMiddleware, middleware.Validate(func() interface{} { return &dto.UpdateHabitRequest{} }))
	r.DELETE("/habit/:habitid", handler.DeleteHabitHandler, middleware.JWTMiddleware, middleware.Validate(func() interface{} { return &dto.UpdateHabitRequest{} }))
}
