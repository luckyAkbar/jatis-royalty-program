package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
)
