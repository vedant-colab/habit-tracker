package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Validate(dtoFactory func() interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Create a fresh DTO per request
			dto := dtoFactory()

			// Bind request body to DTO
			if err := c.Bind(dto); err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input: " + err.Error()})
			}

			// Validate the DTO
			validate := validator.New()
			if err := validate.Struct(dto); err != nil {
				validationErrors := err.(validator.ValidationErrors)

				errors := make(map[string]string)
				for _, fieldErr := range validationErrors {
					errors[fieldErr.Field()] = getErrorMessage(fieldErr)
				}

				return c.JSON(http.StatusBadRequest, echo.Map{"errors": errors})
			}

			// Save the DTO inside context so handler can use it
			c.Set("dto", dto)

			return next(c) // Proceed to next handler
		}
	}
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters long"
	default:
		return fe.Field() + " is invalid"
	}
}
