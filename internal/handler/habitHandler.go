package handler

import (
	"fmt"
	"habit-tracker/internal/dto"
	"habit-tracker/internal/service"
	"habit-tracker/utils"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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

func UpdateHabitHandler(c echo.Context) error {
	claims, ok := c.Get("claims").(jwt.MapClaims)
	fmt.Println(claims)
	if !ok {
		c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid claims",
		})
		return nil
	}

	habitIDStr := c.Param("habitid") // Extract "habitid" from URL

	// Convert to integer if needed
	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid habit ID")
	}

	username := claims["username"].(string)
	updateHabitRequest, err := utils.GetDTO[*dto.UpdateHabitRequest](c)
	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid request context",
		})
		return err
	}
	if err = service.UpdateHabitService(username, habitID, updateHabitRequest); err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
		return err
	}

	c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Habit updated successfully",
	})
	return nil

}

func DeleteHabitHandler(c echo.Context) error {
	claims, ok := c.Get("claims").(jwt.MapClaims)
	fmt.Println(claims)
	if !ok {
		c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid claims",
		})
		return nil
	}

	habitIDStr := c.Param("habitid") // Extract "habitid" from URL

	// Convert to integer if needed
	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid habit ID")
	}

	username := claims["username"].(string)
	updateHabitRequest, err := utils.GetDTO[*dto.UpdateHabitRequest](c)
	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid request context",
		})
		return err
	}
	if err = service.DeleteHabitService(username, habitID, updateHabitRequest); err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
		return err
	}

	c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Habit deleted successfully",
	})
	return nil

}
