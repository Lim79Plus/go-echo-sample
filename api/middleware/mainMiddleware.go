package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetMainMiddleware set main middleware
func SetMainMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
