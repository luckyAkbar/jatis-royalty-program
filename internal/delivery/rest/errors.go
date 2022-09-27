package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInternal           = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	ErrBadRequest         = echo.NewHTTPError(http.StatusBadRequest, "bad request")
	ErrNotFound           = echo.NewHTTPError(http.StatusNotFound, "not found")
	ErrUnauthorized       = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	ErrPreconditionFailed = echo.NewHTTPError(http.StatusPreconditionFailed, "precondition failed")
)
