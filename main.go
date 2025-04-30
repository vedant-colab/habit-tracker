package main

import (
	"habit-tracker/config"
	database "habit-tracker/internal/db"
	"net/http"

	"habit-tracker/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func RegisterRoutes(r *echo.Group) {
	routes.UserRoutes(r)
	routes.HabitRoutes(r)
}

func main() {
	config.LoadEnv()
	database.ConnectDB()

	e := echo.New()
	router := e.Group("/api")
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// Serve Swagger documentation at /swagger/*
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.Use(middleware.CSRF())
	RegisterRoutes(router)

	e.GET("/", func(c echo.Context) error {
		c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"message": "habit-tracker",
		})
		return nil
	})

	e.Logger.Fatal(e.Start(":3000"))
}
