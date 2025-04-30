package handler

import (
	"habit-tracker/internal/dto"
	"habit-tracker/internal/service"
	"habit-tracker/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateHabitHandler(c echo.Context) error {
	habitRequest, err := utils.GetDTO[*dto.CreateHabitRequest](c)
	// return err
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid request context"})
		return err
	}
	err = service.CreateHabitService(habitRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"success": true,
			"error":   err.Error(),
		})
		return err
	}

	c.JSON(http.StatusCreated, echo.Map{
		"success": true,
		"message": "Habit added sucessfully",
	})
	return nil

}
