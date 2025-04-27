package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	ErrDTOTypeAssertion = errors.New("failed to assert DTO to context")
)

func GetDTO[T any](c echo.Context) (T, error) {
	dto, ok := c.Get("dto").(T)
	if !ok {
		var zero T
		return zero, ErrDTOTypeAssertion
	}
	return dto, nil
}
